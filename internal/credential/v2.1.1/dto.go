package credential

import (
	"context"
	"fmt"
	"os"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	dto "github.com/satimoto/go-ocpi/internal/dto/v2.1.1"
)

func (r *CredentialResolver) CreateCredentialDto(ctx context.Context, credential db.Credential) *dto.CredentialDto {
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

	credentialDto := &dto.CredentialDto{
		Token:       util.NilString(credential.ServerToken),
		Url:         &apiDomain,
		PartyID:     &apiPartyID,
		CountryCode: &apiCountryCode,
	}
	credentialDto.BusinessDetail = businessDetailDto

	return credentialDto
}
