package credential

import (
	"net/http"

	"github.com/satimoto/go-ocpi-api/internal/credential"
)

func (r *CredentialResolver) CredentialContextByToken(next http.Handler) http.Handler {
	return credential.CredentialContextByToken(r.Repository, next)
}

func (r *CredentialResolver) CredentialContextByPartyAndCountry(next http.Handler) http.Handler {
	return credential.CredentialContextByPartyAndCountry(r.Repository, next)
}