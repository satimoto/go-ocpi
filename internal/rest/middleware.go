package rest

import (
	"net/http"

	"github.com/satimoto/go-ocpi/internal/middleware"
)

func (rs *RestService) CredentialContextByToken(next http.Handler) http.Handler {
	return middleware.CredentialContextByToken(rs.RepositoryService, next)
}

func (rs *RestService) CredentialContextByPartyAndCountry(next http.Handler) http.Handler {
	return middleware.CredentialContextByPartyAndCountry(rs.RepositoryService, next)
}
