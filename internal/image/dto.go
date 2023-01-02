package image

import (
	"context"

	"github.com/satimoto/go-datastore/pkg/db"
	coreDto "github.com/satimoto/go-ocpi/internal/dto"
)

func (r *ImageResolver) CreateImageDto(ctx context.Context, image db.Image) *coreDto.ImageDto {
	return coreDto.NewImageDto(image)
}

func (r *ImageResolver) CreateImageListDto(ctx context.Context, images []db.Image) []*coreDto.ImageDto {
	list := []*coreDto.ImageDto{}
	
	for _, image := range images {
		list = append(list, r.CreateImageDto(ctx, image))
	}

	return list
}
