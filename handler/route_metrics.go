package handler

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/RyanGuthrie/simple_prom"
	"github.com/prometheus/client_golang/prometheus"
)

var ServerRequestMetrics = RequestMetrics{routeToMetrics: make(map[string]*RouteMetrics)}

type RequestMetrics struct {
	routeToMetrics map[string]*RouteMetrics
	mutex          sync.Mutex
}

func (requestMetrics *RequestMetrics) RouteMetricsFor(id string) *RouteMetrics {
	requestMetrics.mutex.Lock()
	defer requestMetrics.mutex.Unlock()

	requestMetric, found := requestMetrics.routeToMetrics[id]

	if !found {
		requestMetrics.routeToMetrics[id] = NewRouteMetrics(id)
		requestMetric = requestMetrics.routeToMetrics[id]
	}

	return requestMetric
}

type RouteMetrics struct {
	mutex                   sync.Mutex
	Id                      string
	RequestCounter          prometheus.Counter
	RequestLatencies        prometheus.Histogram
	RequestLatenciesSummary prometheus.Summary
	RequestSize             prometheus.Histogram
	ResponseSize            prometheus.Histogram
	statusCodeCounters      map[int]prometheus.Counter
	statusGroupCounters     map[int]prometheus.Counter
}

const kb = 1024
const mb = kb * kb

func NewRouteMetrics(Id string) *RouteMetrics {
	return &RouteMetrics{
		Id: Id,
		RequestCounter: simple_prom.Metrics.NewCounter(prometheus.CounterOpts{
			Namespace: "http",
			Subsystem: Id,
			Name:      "requests_total",
			Help:      "Total number of HTTP requests made for the specific method+route"}),
		statusCodeCounters:  make(map[int]prometheus.Counter),
		statusGroupCounters: make(map[int]prometheus.Counter),
		RequestSize: simple_prom.Metrics.NewHistogram(prometheus.HistogramOpts{
			Namespace: "http",
			Subsystem: Id,
			Name:      "request_bytes",
			Help:      "Bytes per request",
			Buckets: []float64{
				float64(1 * kb),
				float64(5 * kb),
				float64(25 * kb),
				float64(100 * kb),
				float64(500 * kb),
				float64(1 * mb),
				float64(5 * mb)},
		}),
		ResponseSize: simple_prom.Metrics.NewHistogram(prometheus.HistogramOpts{
			Namespace: "http",
			Subsystem: Id,
			Name:      "response_bytes",
			Help:      "Bytes per request",
			Buckets: []float64{
				float64(1 * kb),
				float64(5 * kb),
				float64(25 * kb),
				float64(100 * kb),
				float64(500 * kb),
				float64(1 * mb),
				float64(5 * mb)},
		}),

		RequestLatencies: simple_prom.Metrics.NewHistogram(prometheus.HistogramOpts{
			Namespace: "http",
			Subsystem: Id,
			Name:      "request_latency_ms",
			Help:      "Latency for requests (ms)",
			Buckets: []float64{
				float64(1 * time.Millisecond.Milliseconds()),
				float64(10 * time.Millisecond.Milliseconds()),
				float64(50 * time.Millisecond.Milliseconds()),
				float64(100 * time.Millisecond.Milliseconds()),
				float64(250 * time.Millisecond.Milliseconds()),
				float64(500 * time.Millisecond.Milliseconds()),
				float64(1000 * time.Millisecond.Milliseconds()),
				float64(1500 * time.Millisecond.Milliseconds()),
				float64(2000 * time.Millisecond.Milliseconds()),
				float64(5000 * time.Millisecond.Milliseconds()),
				float64(10000 * time.Millisecond.Milliseconds()),
			},
		}),
		RequestLatenciesSummary: simple_prom.Metrics.NewSummary(prometheus.SummaryOpts{
			Namespace: "http",
			Subsystem: Id,
			Name:      "request_latency_ms_summary",
			Help:      "Latency for requests (ms)",
			Objectives: map[float64]float64{
				.50:   .01,
				.90:   .01,
				.95:   .01,
				.99:   .01,
				.999:  .01,
				.9999: .01,
			},
			// MaxAge:     time.Minute,
			// AgeBuckets: 50000,
			// BufCap:     10000,
		}),
	}
}

func (routeMetrics *RouteMetrics) incrementStatusCodeCounter(statusCode int) {
	routeMetrics.mutex.Lock()
	defer routeMetrics.mutex.Unlock()

	var statusCodeCounter, foundCode = routeMetrics.statusCodeCounters[statusCode]
	var statusGroupCounter, foundGroup = routeMetrics.statusGroupCounters[statusCode/100]

	if !foundCode {
		statusCodeCounter = routeMetrics.newCounterForStatusCode(statusCode)
		routeMetrics.statusCodeCounters[statusCode] = statusCodeCounter
	}

	if !foundGroup {
		statusGroupCounter = routeMetrics.newCounterForStatusGroup(statusCode)
		routeMetrics.statusGroupCounters[statusCode/100] = statusGroupCounter
	}

	statusCodeCounter.Inc()
	statusGroupCounter.Inc()
}

func (routeMetrics *RouteMetrics) newCounterForStatusCode(statusCode int) prometheus.Counter {
	return simple_prom.Metrics.NewCounter(prometheus.CounterOpts{
		Namespace: "http",
		Subsystem: routeMetrics.Id,
		Name:      fmt.Sprintf("statuses:%v", statusCode),
		Help:      fmt.Sprintf("Number of HTTP Status Code %v for the specific method+route", statusCode)})
}

func (routeMetrics *RouteMetrics) newCounterForStatusGroup(statusCode int) prometheus.Counter {
	var group string
	switch {
	case statusCode < 200:
		group = "1XX"
	case statusCode < 300:
		group = "2XX"
	case statusCode < 400:
		group = "3XX"
	case statusCode < 500:
		group = "4XX"
	case statusCode < 600:
		group = "5XX"
	default:
		group = strconv.Itoa(statusCode)
	}

	return simple_prom.Metrics.NewCounter(prometheus.CounterOpts{
		Namespace: "http",
		Subsystem: routeMetrics.Id,
		Name:      fmt.Sprintf("statuses:%v", group),
		Help:      fmt.Sprintf("Number of HTTP Status responses in group %v for the specific method+route", group)})
}
