package api

import (
	"net/http"

	"github.com/satimoto/go-ocpi-api/internal/credential"
)

func (rs *RouterService) CredentialContextByToken(next http.Handler) http.Handler {
	return credential.CredentialContextByToken(rs.RepositoryService, next)
}

func (rs *RouterService) CredentialContextByPartyAndCountry(next http.Handler) http.Handler {
	return credential.CredentialContextByPartyAndCountry(rs.RepositoryService, next)
}
