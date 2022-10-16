package rest

import (
	"github.com/go-chi/chi/v5"
	connector "github.com/satimoto/go-ocpi/internal/connector/v2.1.1"
)

func (rs *RestService) mountConnectors() *chi.Mux {
	connectorResolver := connector.NewResolver(rs.RepositoryService)
	router := chi.NewRouter()

	router.Route("/{connector_id}", func(connectorRouter chi.Router) {
		connectorRouter.Put("/", connectorResolver.UpdateConnector)

		connectorContextRouter := connectorRouter.With(connectorResolver.ConnectorContext(rs.ServiceResolver.SyncService))
		connectorContextRouter.Get("/", connectorResolver.GetConnector)
		connectorContextRouter.Patch("/", connectorResolver.UpdateConnector)
	})

	return router
}
