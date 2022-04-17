package ocpirpc

import (
	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-datastore/util"
)

func NewCreateBusinessDetailParams(input CreateBusinessDetailRequest) db.CreateBusinessDetailParams {
	return db.CreateBusinessDetailParams{
		Name:    input.Name,
		Website: util.SqlNullString(input.Website),
	}
}

func NewCreateImageParams(input CreateImageRequest) db.CreateImageParams {
	return db.CreateImageParams{
		Url:       input.Url,
		Thumbnail: util.SqlNullString(input.Thumbnail),
		Type:      input.Type,
		Category:  db.ImageCategory(input.Category),
		Width:     util.SqlNullInt32(input.Width),
		Height:    util.SqlNullInt32(input.Height),
	}
}
