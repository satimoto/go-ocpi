package credential

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/businessdetail"
	"github.com/satimoto/go-ocpi-api/internal/util"
)

type CredentialPayload struct {
	Token          *string                               `json:"token"`
	Url            *string                               `json:"url"`
	BusinessDetail *businessdetail.BusinessDetailPayload `json:"business_details"`
	PartyID        *string                               `json:"party_id"`
	CountryCode    *string                               `json:"country_code"`
}

func (r *CredentialPayload) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewCredentialPayload(credential db.Credential) *CredentialPayload {
	return &CredentialPayload{
		Token:       util.NilString(credential.ServerToken.String),
		Url:         &credential.Url,
		PartyID:     &credential.PartyID,
		CountryCode: &credential.CountryCode,
	}
}

func NewUpdateCredentialParams(credential db.Credential) db.UpdateCredentialParams {
	return db.UpdateCredentialParams{
		ID:          credential.ID,
		ClientToken: credential.ClientToken,
		ServerToken: credential.ServerToken,
		Url:         credential.Url,
		PartyID:     credential.PartyID,
		CountryCode: credential.CountryCode,
		LastUpdated: credential.LastUpdated,
	}
}

func (r *CredentialResolver) CreateCredentialPayload(ctx context.Context, credential db.Credential) *CredentialPayload {
	apiDomain := os.Getenv("API_DOMAIN")
	apiPartyID := os.Getenv("API_PARTY_ID")
	apiCountryCode := os.Getenv("API_COUNTRY_CODE")
	webDomain := os.Getenv("WEB_DOMAIN")

	imagePayload := r.BusinessDetailResolver.ImageResolver.CreateImagePayload(ctx, db.Image{
		Url:       fmt.Sprintf("%s/logo.png", webDomain),
		Thumbnail: util.SqlNullString(fmt.Sprintf("%s/logo-thumb.png", webDomain)),
		Category:  db.ImageCategoryOPERATOR,
		Type:      "png",
		Width:     util.SqlNullInt32(512),
		Height:    util.SqlNullInt32(512),
	})

	businessDetailPayload := r.BusinessDetailResolver.CreateBusinessDetailPayload(ctx, db.BusinessDetail{
		Name:    "Satimoto",
		Website: util.SqlNullString(webDomain),
	})
	businessDetailPayload.Logo = imagePayload

	credentialPayload := &CredentialPayload{
		Token:       util.NilString(credential.ServerToken.String),
		Url:         &apiDomain,
		PartyID:     &apiPartyID,
		CountryCode: &apiCountryCode,
	}
	credentialPayload.BusinessDetail = businessDetailPayload

	return credentialPayload
}
