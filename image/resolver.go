package image

import (
	"context"
	"database/sql"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/util"
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

func NewImagePayload(image db.Image) *ImagePayload {
	return &ImagePayload{
		Url:       image.Url,
		Thumbnail: util.NilString(image.Thumbnail.String),
		Category:  image.Category,
		Type:      image.Type,
		Width:     util.NilInt32(image.Width.Int32),
		Height:    util.NilInt32(image.Height.Int32),
	}
}

func NewCreateImageParams(payload *ImagePayload) db.CreateImageParams {
	return db.CreateImageParams{
		Url:       payload.Url,
		Thumbnail: util.SqlNullString(payload.Thumbnail),
		Category:  payload.Category,
		Width:     sql.NullInt32{Int32: *payload.Width},
		Height:    sql.NullInt32{Int32: *payload.Height},
	}
}

func (r *ImageResolver) CreateImage(ctx context.Context, payload *ImagePayload) *int64 {
	if payload != nil {
		imageParams := NewCreateImageParams(payload)

		image, err := r.Repository.CreateImage(ctx, imageParams)

		if err != nil {
			return &image.ID
		}
	}

	return nil
}
