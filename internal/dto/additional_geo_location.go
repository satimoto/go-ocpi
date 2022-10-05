package dto

import (
	"net/http"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-ocpi/internal/ocpitype"
)

type AdditionalGeoLocationDto struct {
	Latitude  ocpitype.String `json:"latitude"`
	Longitude ocpitype.String `json:"longitude"`
	Name      *DisplayTextDto `json:"name,omitempty"`
}

func (r *AdditionalGeoLocationDto) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewAdditionalGeoLocationDto(additionalGeoLocation db.AdditionalGeoLocation) *AdditionalGeoLocationDto {
	return &AdditionalGeoLocationDto{
		Latitude:  ocpitype.NewString(additionalGeoLocation.Latitude),
		Longitude: ocpitype.NewString(additionalGeoLocation.Longitude),
	}
}
