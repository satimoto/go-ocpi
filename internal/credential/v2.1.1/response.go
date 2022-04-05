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

type CredentialDto struct {
	Token          *string                           `json:"token"`
	Url            *string                           `json:"url"`
	BusinessDetail *businessdetail.BusinessDetailDto `json:"business_details"`
	CountryCode    *string                           `json:"country_code"`
	PartyID        *string                           `json:"party_id"`
}

func (r *CredentialDto) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewCredentialDto(credential db.Credential) *CredentialDto {
	return &CredentialDto{
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
		CountryCode: credential.CountryCode,
		PartyID:     credential.PartyID,
		IsHub:       credential.IsHub,
		LastUpdated: credential.LastUpdated,
	}
}

func (r *CredentialResolver) CreateCredentialDto(ctx context.Context, credential db.Credential) *CredentialDto {
	apiDomain := os.Getenv("API_DOMAIN")
	apiPartyID := os.Getenv("API_PARTY_ID")
	apiCountryCode := os.Getenv("API_COUNTRY_CODE")
	webDomain := os.Getenv("WEB_DOMAIN")

	imageDto := r.BusinessDetailResolver.ImageResolver.CreateImageDto(ctx, db.Image{
		Url:       fmt.Sprintf("%s/logo.png", webDomain),
		Thumbnail: util.SqlNullString(fmt.Sprintf("%s/logo-thumb.png", webDomain)),
		Category:  db.ImageCategoryOPERATOR,
		Type:      "png",
		Width:     util.SqlNullInt32(512),
		Height:    util.SqlNullInt32(512),
	})

	businessDetailDto := r.BusinessDetailResolver.CreateBusinessDetailDto(ctx, db.BusinessDetail{
		Name:    "Satimoto",
		Website: util.SqlNullString(webDomain),
	})
	businessDetailDto.Logo = imageDto

	credentialDto := &CredentialDto{
		Token:       util.NilString(credential.ServerToken.String),
		Url:         &apiDomain,
		PartyID:     &apiPartyID,
		CountryCode: &apiCountryCode,
	}
	credentialDto.BusinessDetail = businessDetailDto

	return credentialDto
}
