package dto

import (
	"net/http"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
)

type BusinessDetailDto struct {
	Name    string          `json:"name"`
	Website *string         `json:"website,omitempty"`
	Logo    *ImageDto `json:"logo,omitempty"`
}

func (r *BusinessDetailDto) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewBusinessDetailDto(businessDetail db.BusinessDetail) *BusinessDetailDto {
	return &BusinessDetailDto{
		Name:    businessDetail.Name,
		Website: util.NilString(businessDetail.Website),
	}
}
