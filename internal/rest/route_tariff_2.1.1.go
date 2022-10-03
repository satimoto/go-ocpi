package rest

import (
	"github.com/go-chi/chi/v5"
	tariff "github.com/satimoto/go-ocpi/internal/tariff/v2.1.1"
	"github.com/satimoto/go-ocpi/internal/version"
)

func (rs *RestService) mountTariffs() *chi.Mux {
	tariffResolver := tariff.NewResolver(rs.RepositoryService)
	rs.SyncService.AddHandler(version.VERSION_2_1_1, tariffResolver)

	router := chi.NewRouter()
	router.Use(rs.CredentialContextByToken)

	router.Route("/{country_code}/{party_id}/{tariff_id}", func(tariffRouter chi.Router) {
		tariffRouter.Put("/", tariffResolver.UpdateTariff)

		tariffContextRouter := tariffRouter.With(tariffResolver.TariffContext(rs.SyncService))
		tariffContextRouter.Get("/", tariffResolver.GetTariff)
		tariffContextRouter.Delete("/", tariffResolver.DeleteTariff)
		tariffContextRouter.Patch("/", tariffResolver.UpdateTariff)
	})

	return router
}
