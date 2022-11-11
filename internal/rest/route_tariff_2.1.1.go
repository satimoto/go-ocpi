package rest

import (
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	coreTariff "github.com/satimoto/go-ocpi/internal/tariff"
	tariff "github.com/satimoto/go-ocpi/internal/tariff/v2.1.1"
	"github.com/satimoto/go-ocpi/internal/version"
)

func (rs *RestService) mountTariffs() *chi.Mux {
	tariffResolver := tariff.NewResolver(rs.RepositoryService, rs.ServiceResolver)
	rs.ServiceResolver.SyncService.AddHandler(version.VERSION_2_1_1, coreTariff.IDENTIFIER, tariffResolver)

	router := chi.NewRouter()
	router.Use(middleware.Timeout(30 * time.Second))
	router.Use(rs.CredentialContextByToken)

	router.Route("/{country_code}/{party_id}/{tariff_id}", func(tariffRouter chi.Router) {
		tariffRouter.Put("/", tariffResolver.UpdateTariff)

		tariffContextRouter := tariffRouter.With(tariffResolver.TariffContext(rs.ServiceResolver.SyncService))
		tariffContextRouter.Get("/", tariffResolver.GetTariff)
		tariffContextRouter.Delete("/", tariffResolver.DeleteTariff)
		tariffContextRouter.Patch("/", tariffResolver.UpdateTariff)
	})

	return router
}
