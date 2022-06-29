package credential

import (
	"net/http"

	"github.com/satimoto/go-ocpi/internal/middleware"
)

func (r *CredentialResolver) CredentialContextByToken(next http.Handler) http.Handler {
	return middleware.CredentialContextByToken(r.Repository, next)
}

func (r *CredentialResolver) CredentialContextByPartyAndCountry(next http.Handler) http.Handler {
	return middleware.CredentialContextByPartyAndCountry(r.Repository, next)
}
