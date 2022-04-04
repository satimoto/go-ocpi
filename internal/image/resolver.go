package image

import (
	"context"
	"database/sql"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/util"
)

type ImageRepository interface {
	CreateImage(ctx context.Context, arg db.CreateImageParams) (db.Image, error)
	GetImage(ctx context.Context, id int64) (db.Image, error)
}

type ImageResolver struct {
	Repository ImageRepository
}

func NewResolver(repositoryService *db.RepositoryService) *ImageResolver {
	repo := ImageRepository(repositoryService)
	return &ImageResolver{repo}
}

func NewImageDto(image db.Image) *ImageDto {
	return &ImageDto{
		Url:       image.Url,
		Thumbnail: util.NilString(image.Thumbnail.String),
		Category:  image.Category,
		Type:      image.Type,
		Width:     util.NilInt32(image.Width.Int32),
		Height:    util.NilInt32(image.Height.Int32),
	}
}

func NewCreateImageParams(dto *ImageDto) db.CreateImageParams {
	return db.CreateImageParams{
		Url:       dto.Url,
		Thumbnail: util.SqlNullString(dto.Thumbnail),
		Category:  dto.Category,
		Width:     sql.NullInt32{Int32: *dto.Width},
		Height:    sql.NullInt32{Int32: *dto.Height},
	}
}

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
