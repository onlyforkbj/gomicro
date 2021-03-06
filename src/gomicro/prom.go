package main

import (
	"net/http"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

var http_requests_total = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Namespace: "go",
		Subsystem: "micro",
		Name:      "http_requests_total",
		Help:      "The number of HTTP requests.",
	},
	[]string{"role", "method", "path", "status"},
)

var http_requests_latency = prometheus.NewSummaryVec(
	prometheus.SummaryOpts{
		Namespace: "go",
		Subsystem: "micro",
		Name:      "http_request_latency",
		Help:      "The latency of HTTP requests.",
	},
	[]string{"method", "path", "status"},
)

var http_healthcheck_total = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Namespace: "go",
		Subsystem: "micro",
		Name:      "http_healthcheck_total",
		Help:      "The number of healthcheck requests.",
	},
	[]string{"target", "status"},
)

var http_healthcheck_latency = prometheus.NewSummaryVec(
	prometheus.SummaryOpts{
		Namespace: "go",
		Subsystem: "micro",
		Name:      "http_healthcheck_latency",
		Help:      "The latency of healthcheck requests.",
	},
	[]string{"target"},
)

func PromRegister() {
	prometheus.MustRegister(http_requests_total)
	//prometheus.MustRegister(http_requests_latency)

	prometheus.MustRegister(http_healthcheck_total)
	prometheus.MustRegister(http_healthcheck_latency)
}

func isWSRequest(req *http.Request) bool {
	return strings.ToLower(req.Header.Get("Upgrade")) == "websocket" &&
		strings.ToLower(req.Header.Get("Connection")) == "upgrade"
}

func isPromRequest(req *http.Request) bool {
	return strings.Contains(strings.ToLower(req.URL.Path), "/metrics")
}
