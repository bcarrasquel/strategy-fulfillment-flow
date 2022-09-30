package metrics

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// CreateCounter prometheus
type CounterInterface interface {
	WithLabelValues(lvs ...string) prometheus.Counter
}

var metrics = make(map[string]CounterInterface)
var serverListen = http.ListenAndServe
var newCounter = prometheus.NewCounterVec
var register = prometheus.Register

func InitMetrics() map[string]CounterInterface {
	metrics["read"] = createCounterVec("fo_event_read_metric", "Number of total events received from queue clusters")
	metrics["process"] = createCounterVec("fo_event_process_metric", "Number of total events process from queue clusters")
	metrics["publish"] = createCounterVec("fo_event_public_metric", "Number of total events sent to queue")
	metrics["error"] = createCounterVec("fo_event_error_metric", "Number of total events with error")
	return metrics
}

func Inc(metric CounterInterface, labels ...string) {
	metric.WithLabelValues(labels...).Inc()
}

func createCounterVec(name string, help string) CounterInterface {
	var options = prometheus.CounterOpts{
		Name: name,
		Help: help,
	}

	var counter = newCounter(options, []string{"topic", "error"})

	if err := register(counter); err != nil {
		fmt.Println("Register metrics error")
	}

	return counter
}

// Expose start http server for metrics
func Expose() error {
	http.Handle("/metrics", promhttp.Handler())
	return serverListen(":7766", nil)
}
