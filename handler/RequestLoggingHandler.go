package handler

import (
	"fmt"
	"github.com/felixge/httpsnoop"
	"log"
	"net/http"
	"strings"
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

func RequestLoggingHandler(handler http.Handler) http.Handler {
	overallRouteMetrics := NewRouteMetrics("overall")

	return http.HandlerFunc(
		func(response http.ResponseWriter, request *http.Request) {
			id := idOf(request)
			routeMetrics := ServerRequestMetrics.RouteMetricsFor(id)

			routeMetrics.RequestCounter.Inc()
			overallRouteMetrics.RequestCounter.Inc()

			m := httpsnoop.CaptureMetrics(handler, response, request)

			path := request.URL.Path
			responseStatus := m.Code
			responseSize := m.Written
			duration := m.Duration

			routeMetrics.incrementStatusCodeCounter(responseStatus)
			overallRouteMetrics.incrementStatusCodeCounter(responseStatus)

			routeMetrics.RequestLatencies.Observe(duration.Seconds())
			overallRouteMetrics.RequestLatencies.Observe(duration.Seconds())

			log.Printf(
				"[%v %v], Status: [%v], Size: [%v], Time: [%v]",
				request.Method,
				path,
				responseStatus,
				responseSize,
				duration,
			)

		})
}
