package rest

import (
	"github.com/go-chi/chi/v5"
	credential "github.com/satimoto/go-ocpi/internal/credential/v2.1.1"
)

func (rs *RestService) mountCredentials() *chi.Mux {
	credentialResolver := credential.NewResolver(rs.RepositoryService, rs.SyncService, rs.OcpiRequester)

	router := chi.NewRouter()
	router.Delete("/", credentialResolver.DeleteCredential)
	router.Post("/", credentialResolver.UpdateCredential)
	router.Put("/", credentialResolver.UpdateCredential)

	credentialContextRouter := router.With(credentialResolver.CredentialContextByToken)
	credentialContextRouter.Get("/", credentialResolver.GetCredential)

	return router
}
