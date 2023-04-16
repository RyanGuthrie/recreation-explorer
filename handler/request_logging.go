package handler

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/felixge/httpsnoop"
	"github.com/julienschmidt/httprouter"
)

func idOf(request *http.Request) string {
	urlPath := request.URL.Path[1:] // Trim leading slash
	sanitizedURL := strings.ReplaceAll(urlPath, "/", "_")
	sanitizedURL = strings.ReplaceAll(sanitizedURL, ".", "_")
	if sanitizedURL == "" {
		sanitizedURL = "root"
	}

	return fmt.Sprintf("%v_%v", request.Method, sanitizedURL)
}

type RequestLoggingHandle struct {
	overallMetrics *RouteMetrics
}

func NewRequestLoggingHandler() RequestLoggingHandle {
	return RequestLoggingHandle{
		overallMetrics: NewRouteMetrics("overall"),
	}
}

func (h RequestLoggingHandle) ToHandle(handler http.Handler) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		handler.ServeHTTP(writer, request)
	}
}

func (h RequestLoggingHandle) Handle(handler httprouter.Handle) httprouter.Handle {
	return func(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
		startTime := time.Now()

		id := idOf(request)
		routeMetrics := ServerRequestMetrics.RouteMetricsFor(id)

		routeMetrics.RequestCounter.Inc()
		h.overallMetrics.RequestCounter.Inc()

		handlerAdapter := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			handler(writer, request, params)
		})

		requestSize := request.ContentLength
		m := httpsnoop.CaptureMetrics(handlerAdapter, response, request)

		routeMetrics.incrementStatusCodeCounter(m.Code)
		h.overallMetrics.incrementStatusCodeCounter(m.Code)

		routeMetrics.RequestLatencies.Observe(float64(m.Duration.Milliseconds()))
		h.overallMetrics.RequestLatencies.Observe(float64(m.Duration.Milliseconds()))
		routeMetrics.RequestLatenciesSummary.Observe(float64(m.Duration.Milliseconds()))
		h.overallMetrics.RequestLatenciesSummary.Observe(float64(m.Duration.Milliseconds()))

		routeMetrics.RequestSize.Observe(float64(requestSize))
		routeMetrics.ResponseSize.Observe(float64(m.Written))

		endTime := time.Now()

		/**
		  GET /metrics
		    Start: [2023-04-02T22:27:31Z] Proto:  [1.1] RequestBody:  [0]
		    End:   [2023-04-02T22:27:31Z] Status: [200] ResponseBody: [4749] Duration [3.399459ms]
		*/
		log.Printf("%v %v\n"+
			"  Start: [%v] Proto:  [%v] RequestBody:  [%v]\n"+
			"  End:   [%v] Status: [%v] ResponseBody: [%v] Duration [%v]\n",
			request.Method,
			request.URL.Path,
			startTime.UTC().Format(time.RFC3339),
			fmt.Sprintf("%d.%d", request.ProtoMajor, request.ProtoMinor),
			requestSize,
			endTime.UTC().Format(time.RFC3339),
			m.Code,
			m.Written,
			m.Duration,
		)
	}
}
