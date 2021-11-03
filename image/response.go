package image

import (
	"context"
	"net/http"

	"github.com/satimoto/go-datastore/db"
)


type ImagePayload struct {
	Url       string           `json:"url"`
	Thumbnail *string          `json:"thumbnail,omitempty"`
	Category  db.ImageCategory `json:"category"`
	Type      string           `json:"type"`
	Width     *int32           `json:"width,omitempty"`
	Height    *int32           `json:"height,omitempty"`
}

func (r *ImagePayload) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func (r *ImageResolver) CreateImagePayload(ctx context.Context, image db.Image) *ImagePayload {
	return NewImagePayload(image)
}

func (r *ImageResolver) CreateImageListPayload(ctx context.Context, images []db.Image) []*ImagePayload {
	list := []*ImagePayload{}
	for _, image := range images {
		list = append(list, r.CreateImagePayload(ctx, image))
	}
	return list
}
