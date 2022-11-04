package image

import (
	"context"
	"log"

	coreDto "github.com/satimoto/go-ocpi/internal/dto"
	metrics "github.com/satimoto/go-ocpi/internal/metric"
)

func (r *ImageResolver) CreateImage(ctx context.Context, imageDto *coreDto.ImageDto) *int64 {
	if imageDto != nil {
		imageParams := NewCreateImageParams(imageDto)
		image, err := r.Repository.CreateImage(ctx, imageParams)

		if err != nil {
			metrics.RecordError("OCPI116", "Error creating image", err)
			log.Printf("OCPI116: Params=%#v", imageParams)
			return nil
		}

		return &image.ID
	}

	return nil
}
