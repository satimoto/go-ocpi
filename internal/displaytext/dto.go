package displaytext

import (
	"context"
	"net/http"

	"github.com/satimoto/go-datastore/db"
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

func (r *DisplayTextResolver) CreateDisplayTextDto(ctx context.Context, displayText db.DisplayText) *DisplayTextDto {
	return NewDisplayTextDto(displayText)
}

func (r *DisplayTextResolver) CreateDisplayTextListDto(ctx context.Context, displayTexts []db.DisplayText) []*DisplayTextDto {
	list := []*DisplayTextDto{}
	for _, displayText := range displayTexts {
		list = append(list, r.CreateDisplayTextDto(ctx, displayText))
	}
	return list
}
