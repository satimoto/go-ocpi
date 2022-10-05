package dto

import (
	"net/http"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	coreDto "github.com/satimoto/go-ocpi/internal/dto"
	"github.com/satimoto/go-ocpi/internal/ocpitype"
)

type OcpiCredentialDto struct {
	Data          *CredentialDto `json:"data,omitempty"`
	StatusCode    int16          `json:"status_code"`
	StatusMessage string         `json:"status_message"`
	Timestamp     ocpitype.Time  `json:"timestamp"`
}

type CredentialDto struct {
	Token          *string                    `json:"token"`
	Url            *string                    `json:"url"`
	BusinessDetail *coreDto.BusinessDetailDto `json:"business_details"`
	CountryCode    *string                    `json:"country_code"`
	PartyID        *string                    `json:"party_id"`
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
