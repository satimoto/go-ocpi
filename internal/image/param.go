package image

import (
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	coreDto "github.com/satimoto/go-ocpi/internal/dto"
)

func NewCreateImageParams(imageDto *coreDto.ImageDto) db.CreateImageParams {
	return db.CreateImageParams{
		Url:       imageDto.Url,
		Thumbnail: util.SqlNullString(imageDto.Thumbnail),
		Category:  imageDto.Category,
		Width:     util.SqlNullInt32(imageDto.Width),
		Height:    util.SqlNullInt32(imageDto.Height),
	}
}
