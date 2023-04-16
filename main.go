package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/RyanGuthrie/simple_prom"
	"github.com/julienschmidt/httprouter"
	"github.com/prometheus/client_golang/prometheus"
	"server/controller"
	"server/domain"
	"server/handler"
	"server/prompt"
)

var interactive = flag.Bool("i", false, "If specified, starts in interactive mode using the CLI")

func init() {
	log.SetFlags(0)
	flag.Parse()
}

func main() {
	if *interactive {
		var cursorPos int = 0
		for {
			state, err := domain.GetState(&cursorPos)

			if err == prompt.ExitError {
				log.Println("Exiting, goodbye!")
				return
			} else if err != nil {
				log.Fatalln(err)
			}

			if err = state.Explore(); err != nil {
				log.Fatalln(err)
			}
		}
	} else {
		startTimeSeconds := time.Now().Unix()

		closeChan := simple_prom.Metrics.NewGaugeFunc(
			prometheus.GaugeOpts{
				Namespace: "server",
				Name:      "uptime_seconds",
				Help:      "Uptime of server in seconds",
			},
			15*time.Second,
			func() float64 {
				return float64(time.Now().Unix() - startTimeSeconds)
			})

		startHttpServer()

		close(closeChan)
	}
}

func startHttpServer() {
	requestLoggingHandler := handler.NewRequestLoggingHandler()

	routes := []controller.Route{
		{Verb: http.MethodGet, Path: "/metrics", Handler: requestLoggingHandler.ToHandle(simple_prom.Metrics.Handler)},
		{Verb: http.MethodGet, Path: "/state", Handler: controller.StateIndex},
		{Verb: http.MethodGet, Path: "/state/:state/facility", Handler: controller.FacilityIndex},
	}

	router := httprouter.New()
	for _, route := range routes {
		switch route.Verb {
		case http.MethodGet:
			router.GET(route.Path, requestLoggingHandler.Handle(route.Handler))
		case http.MethodPut:
			router.PUT(route.Path, requestLoggingHandler.Handle(route.Handler))
		case http.MethodPost:
			router.POST(route.Path, requestLoggingHandler.Handle(route.Handler))
		case http.MethodDelete:
			router.DELETE(route.Path, requestLoggingHandler.Handle(route.Handler))
		case http.MethodHead:
			router.HEAD(route.Path, requestLoggingHandler.Handle(route.Handler))
		case http.MethodOptions:
			router.OPTIONS(route.Path, requestLoggingHandler.Handle(route.Handler))
		case http.MethodPatch:
			router.PATCH(route.Path, requestLoggingHandler.Handle(route.Handler))
		default:
			log.Fatalf("Unsupported verb: %v when constructing route for %v\n", route.Verb, route.Path)
		}
	}

	router.GET("/", requestLoggingHandler.Handle(controller.NewIndex(routes)))

	if err := http.ListenAndServe(":8000", router); err != nil {
		log.Panicf("Failed starting HTTP server: %v", err)
	}

	log.Println("Exiting")
}
