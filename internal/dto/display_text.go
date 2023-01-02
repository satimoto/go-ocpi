package dto

import (
	"net/http"

	"github.com/satimoto/go-datastore/pkg/db"
)

type DisplayTextDto struct {
	Language string `json:"language"`
	Text     string `json:"text"`
}

func (r *DisplayTextDto) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewDisplayTextDto(displayText db.DisplayText) *DisplayTextDto {
	return &DisplayTextDto{
		Language: displayText.Language,
		Text:     displayText.Text,
	}
}
