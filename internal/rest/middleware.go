package rest

import (
	"net/http"

	"github.com/satimoto/go-ocpi-api/internal/credential"
)

func (rs *RestService) CredentialContextByToken(next http.Handler) http.Handler {
	return credential.CredentialContextByToken(rs.RepositoryService, next)
}

func (rs *RestService) CredentialContextByPartyAndCountry(next http.Handler) http.Handler {
	return credential.CredentialContextByPartyAndCountry(rs.RepositoryService, next)
}
