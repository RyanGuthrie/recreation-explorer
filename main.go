package main

import (
	"flag"
	"fmt"
	"github.com/RyanGuthrie/simple_prom"
	"log"
	"net/http"
	"server/controller"
	"server/domain"
	"server/handler"
	"server/prompt"
	"time"
)
import "golang.org/x/exp/maps"

var interactive = flag.Bool("i", false, "If specified, starts in interactive mode using the CLI")

func init() {
	log.SetFlags(0)
	flag.Parse()
}

func main() {
	fmt.Printf("Value of interactive: %v\n", *interactive)

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
		startHttpServer()
	}
}

func startHttpServer() {
	routes := make(map[string]func(http.ResponseWriter, *http.Request))
	routes["/states"] = controller.StateIndex

	mux := &http.ServeMux{}

	for route, handlerFunc := range routes {
		mux.HandleFunc(route, handlerFunc)
	}

	mux.HandleFunc("/", controller.NewIndex(maps.Keys(routes)))
	mux.Handle("/metrics", simple_prom.Metrics.Handler)

	srv := &http.Server{
		Addr:         ":8000",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  30 * time.Second,
		Handler:      handler.RequestLoggingHandler(mux),
	}

	err := srv.ListenAndServe()

	if err != nil {
		log.Panicf("Failed starting HTTP server: %v", err)
	}

	log.Println("Exiting")
}
