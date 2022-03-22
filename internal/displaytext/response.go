package displaytext

import (
	"context"
	"net/http"

	"github.com/satimoto/go-datastore/db"
)


type DisplayTextPayload struct {
	Language string `json:"language"`
	Text     string `json:"text"`
}

func (r *DisplayTextPayload) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewDisplayTextPayload(displayText db.DisplayText) *DisplayTextPayload {
	return &DisplayTextPayload{
		Language: displayText.Language,
		Text:     displayText.Text,
	}
}

func NewCreateDisplayTextParams(payload *DisplayTextPayload) db.CreateDisplayTextParams {
	return db.CreateDisplayTextParams{
		Language: payload.Language,
		Text:     payload.Text,
	}
}

func (r *DisplayTextResolver) CreateDisplayTextPayload(ctx context.Context, displayText db.DisplayText) *DisplayTextPayload {
	return NewDisplayTextPayload(displayText)
}

func (r *DisplayTextResolver) CreateDisplayTextListPayload(ctx context.Context, displayTexts []db.DisplayText) []*DisplayTextPayload {
	list := []*DisplayTextPayload{}
	for _, displayText := range displayTexts {
		list = append(list, r.CreateDisplayTextPayload(ctx, displayText))
	}
	return list
}
