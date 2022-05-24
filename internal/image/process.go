package image

import (
	"context"
	"log"

	"github.com/satimoto/go-datastore/pkg/util"
)

func (r *ImageResolver) CreateImage(ctx context.Context, dto *ImageDto) *int64 {
	if dto != nil {
		imageParams := NewCreateImageParams(dto)
		image, err := r.Repository.CreateImage(ctx, imageParams)
		
		if err != nil {
			util.LogOnError("OCPI116", "Error creating image", err)
			log.Printf("OCPI116: Params=%#v", imageParams)
			return nil
		}
			
		return &image.ID
	}

	return nil
}
