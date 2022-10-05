package dto

import (
	"net/http"

	"github.com/satimoto/go-datastore/pkg/db"
)

type ElementDto struct {
	PriceComponents []*PriceComponentDto `json:"price_components"`
	Restrictions    *ElementRestrictionDto      `json:"restrictions,omitempty"`
}

func (r *ElementDto) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewElementDto(element db.Element) *ElementDto {
	return &ElementDto{}
}
