package businessdetail_test

import (
	"context"
	"database/sql"
	"encoding/json"
	"testing"

	dbMocks "github.com/satimoto/go-datastore-mocks/db"
	"github.com/satimoto/go-datastore/db"
	businessDetailMocks "github.com/satimoto/go-ocpi-api/internal/businessdetail/mocks"
	"github.com/satimoto/go-ocpi-api/internal/util"
	"github.com/satimoto/go-ocpi-api/test/mocks"
)

func TestCreateBusinessDetailDto(t *testing.T) {
	ctx := context.Background()

	t.Run("Name only", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		businessDetailResolver := businessDetailMocks.NewResolver(mockRepository)

		businessDetail := db.BusinessDetail{
			Name: "Business Name",
		}
		response := businessDetailResolver.CreateBusinessDetailDto(ctx, businessDetail)
		responseJson, _ := json.Marshal(response)

		mocks.CompareJson(t, responseJson, []byte(`{
			"name": "Business Name"
		}`))
	})

	t.Run("Name and Website", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		businessDetailResolver := businessDetailMocks.NewResolver(mockRepository)

		businessDetail := db.BusinessDetail{
			Name:    "Business Name",
			Website: util.SqlNullString("https://business.com"),
		}
		response := businessDetailResolver.CreateBusinessDetailDto(ctx, businessDetail)
		responseJson, _ := json.Marshal(response)

		mocks.CompareJson(t, responseJson, []byte(`{
			"name": "Business Name",
			"website": "https://business.com"
		}`))
	})

	t.Run("Name, Website and Logo", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		businessDetailResolver := businessDetailMocks.NewResolver(mockRepository)

		businessDetail := db.BusinessDetail{
			Name:    "Business Name",
			Website: util.SqlNullString("https://business.com"),
			LogoID:  sql.NullInt64{Int64: 1, Valid: true},
		}
		image := db.Image{
			Url:      "https://business.com/logo.png",
			Category: db.ImageCategoryOPERATOR,
			Type:     "png",
		}

		mockRepository.SetGetImageMockData(dbMocks.ImageMockData{Image: image, Error: nil})

		response := businessDetailResolver.CreateBusinessDetailDto(ctx, businessDetail)
		responseJson, _ := json.Marshal(response)

		mocks.CompareJson(t, responseJson, []byte(`{
			"name": "Business Name",
			"website": "https://business.com",
			"logo": {
				"url": "https://business.com/logo.png",
				"category": "OPERATOR",
				"type": "png"
			}
		}`))
	})
}
