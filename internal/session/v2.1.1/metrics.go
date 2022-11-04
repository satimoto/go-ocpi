package session

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/satimoto/go-datastore/pkg/db"
)

var (
	metricSessionsActive = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "ocpi_sessions_active",
		Help: "The number of active sessions",
	})
	metricSessionsCompleted = promauto.NewCounter(prometheus.CounterOpts{
		Name: "ocpi_sessions_completed_total",
		Help: "The number of completed sessions",
	})
	metricSessionsInvalid = promauto.NewCounter(prometheus.CounterOpts{
		Name: "ocpi_sessions_invalid_total",
		Help: "The number of invalid sessions",
	})
	metricSessionsTotal = promauto.NewCounter(prometheus.CounterOpts{
		Name: "ocpi_sessions_total",
		Help: "The total number of sessions",
	})
)

func (r *SessionResolver) updateMetrics(session db.Session, sessionCreated bool) {
	if sessionCreated {
		metricSessionsActive.Inc()
		metricSessionsTotal.Inc()
	}

	if session.Status == db.SessionStatusTypeCOMPLETED {
		metricSessionsCompleted.Inc()
		metricSessionsActive.Dec()
	}

	if session.Status == db.SessionStatusTypeINVALID {
		metricSessionsInvalid.Inc()
		metricSessionsActive.Dec()
	}
}
