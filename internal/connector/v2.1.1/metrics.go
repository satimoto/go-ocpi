package connector

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	metricConnectorsTotal = promauto.NewCounter(prometheus.CounterOpts{
		Name: "ocpi_connectors_total",
		Help: "The total number of connectors",
	})
)
