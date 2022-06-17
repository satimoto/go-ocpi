package credential

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	"github.com/satimoto/go-ocpi-api/internal/businessdetail"
	"github.com/satimoto/go-ocpi-api/internal/ocpitype"
)

type OcpiCredentialDto struct {
	Data          *CredentialDto `json:"data,omitempty"`
	StatusCode    int16          `json:"status_code"`
	StatusMessage string         `json:"status_message"`
	Timestamp     ocpitype.Time  `json:"timestamp"`
}

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
		Token:       util.NilString(credential.ServerToken),
		Url:         &credential.Url,
		PartyID:     &credential.PartyID,
		CountryCode: &credential.CountryCode,
	}
}

func (r *CredentialResolver) CreateCredentialDto(ctx context.Context, credential db.Credential) *CredentialDto {
	apiDomain := os.Getenv("API_DOMAIN")
	apiPartyID := os.Getenv("PARTY_ID")
	apiCountryCode := os.Getenv("COUNTRY_CODE")
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
		Token:       util.NilString(credential.ServerToken),
		Url:         &apiDomain,
		PartyID:     &apiPartyID,
		CountryCode: &apiCountryCode,
	}
	credentialDto.BusinessDetail = businessDetailDto

	return credentialDto
}
