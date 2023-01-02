package cdr

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	metricCdrsTotal = promauto.NewCounter(prometheus.CounterOpts{
		Name: "ocpi_cdrs_total",
		Help: "The total number of cdrs",
	})
)
