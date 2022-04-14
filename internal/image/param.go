package image

import (
	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-datastore/util"
)

func NewCreateImageParams(dto *ImageDto) db.CreateImageParams {
	return db.CreateImageParams{
		Url:       dto.Url,
		Thumbnail: util.SqlNullString(dto.Thumbnail),
		Category:  dto.Category,
		Width:     util.SqlNullInt32(dto.Width),
		Height:    util.SqlNullInt32(dto.Height),
	}
}
