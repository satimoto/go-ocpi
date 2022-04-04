package versiondetail_test

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	dbMocks "github.com/satimoto/go-datastore-mocks/db"
	versionDetailMocks "github.com/satimoto/go-ocpi-api/internal/versiondetail/v2.1.1/mocks"
	"github.com/satimoto/go-ocpi-api/test/mocks"
)

func TestCreateVersionDetailDto(t *testing.T) {
	ctx := context.Background()

	os.Setenv("API_DOMAIN", "http://localhost:8080")
	defer os.Unsetenv("API_DOMAIN")

	t.Run("Create dto", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		versionDetailResolver := versionDetailMocks.NewResolver(mockRepository, mocks.NewOCPIRequester(mockHTTPRequester))

		response := versionDetailResolver.CreateVersionDetailDto(ctx)
		responseJson, _ := json.Marshal(response)

		mocks.CompareJson(t, responseJson, []byte(`{
			"version": "2.1.1",
			"endpoints": [{
				"identifier": "locations",
				"url": "http://localhost:8080/2.1.1/locations"
			}]
		}`))
	})
}
