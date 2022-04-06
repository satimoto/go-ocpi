package image

import (
	"context"
)

func (r *ImageResolver) CreateImage(ctx context.Context, dto *ImageDto) *int64 {
	if dto != nil {
		imageParams := NewCreateImageParams(dto)

		image, err := r.Repository.CreateImage(ctx, imageParams)

		if err != nil {
			return &image.ID
		}
	}

	return nil
}
