package ocpi

import (
	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/ocpirpc"
)

func NewBusinessDetailResponse(businessDetail db.BusinessDetail) *ocpirpc.BusinessDetailResponse {
	return &ocpirpc.BusinessDetailResponse{
		Name:    businessDetail.Name,
		Website: businessDetail.Website.String,
	}
}

func NewImageResponse(image db.Image) *ocpirpc.ImageResponse {
	return &ocpirpc.ImageResponse{
		Url:       image.Url,
		Thumbnail: image.Thumbnail.String,
		Category:  string(image.Category),
		Type:      image.Type,
		Width:     image.Width.Int32,
		Height:    image.Height.Int32,
	}
}
