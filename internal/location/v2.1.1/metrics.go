package location

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	metricLocationsTotal = promauto.NewCounter(prometheus.CounterOpts{
		Name: "ocpi_locations_total",
		Help: "The total number of locations",
	})
)
