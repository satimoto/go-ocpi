package dto

import (
	"net/http"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
)

type ImageDto struct {
	Url       string           `json:"url"`
	Thumbnail *string          `json:"thumbnail,omitempty"`
	Category  db.ImageCategory `json:"category"`
	Type      string           `json:"type"`
	Width     *int32           `json:"width,omitempty"`
	Height    *int32           `json:"height,omitempty"`
}

func (r *ImageDto) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewImageDto(image db.Image) *ImageDto {
	return &ImageDto{
		Url:       image.Url,
		Thumbnail: util.NilString(image.Thumbnail),
		Category:  image.Category,
		Type:      image.Type,
		Width:     util.NilInt32(image.Width.Int32),
		Height:    util.NilInt32(image.Height.Int32),
	}
}
