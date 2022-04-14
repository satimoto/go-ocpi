package ocpirpc

import "github.com/satimoto/go-datastore/db"

func NewBusinessDetailResponse(businessDetail db.BusinessDetail) *BusinessDetailResponse {
	return &BusinessDetailResponse{
		Name:    businessDetail.Name,
		Website: businessDetail.Website.String,
	}
}

func NewImageResponse(image db.Image) *ImageResponse {
	return &ImageResponse{
		Url:       image.Url,
		Thumbnail: image.Thumbnail.String,
		Category:  string(image.Category),
		Type:      image.Type,
		Width:     image.Width.Int32,
		Height:    image.Height.Int32,
	}
}
