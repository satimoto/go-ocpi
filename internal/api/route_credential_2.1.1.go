package api

import (
	"github.com/go-chi/chi/v5"
	credential "github.com/satimoto/go-ocpi-api/internal/credential/v2.1.1"
)

func (rs *RouterService) mountCredentials() *chi.Mux {
	credentialResolver := credential.NewResolver(rs.RepositoryService)
	router := chi.NewRouter()

	router.Delete("/", credentialResolver.DeleteCredential)
	router.Post("/", credentialResolver.UpdateCredential)
	router.Put("/", credentialResolver.UpdateCredential)

	credentialContextRouter := router.With(credentialResolver.CredentialContextByToken)
	credentialContextRouter.Get("/", credentialResolver.GetCredential)

	return router
}
