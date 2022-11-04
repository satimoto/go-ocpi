package evse

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	metricEvsesTotal = promauto.NewCounter(prometheus.CounterOpts{
		Name: "ocpi_evses_total",
		Help: "The total number of evses",
	})
	metricEvsesStatus = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "ocpi_evses_status_total",
		Help: "The total number of evses",
	}, []string{"status"})
)
