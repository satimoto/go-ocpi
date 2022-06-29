package versiondetail_test

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	dbMocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	transportationMocks "github.com/satimoto/go-ocpi/internal/transportation/mocks"
	versionDetailMocks "github.com/satimoto/go-ocpi/internal/versiondetail/mocks"
	"github.com/satimoto/go-ocpi/test/mocks"
)

func TestCreateVersionDetailDto(t *testing.T) {
	ctx := context.Background()

	os.Setenv("API_DOMAIN", "http://localhost:8080")
	defer os.Unsetenv("API_DOMAIN")

	t.Run("Create dto", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		versionDetailResolver := versionDetailMocks.NewResolverWithServices(mockRepository, transportationMocks.NewOcpiRequester(mockHTTPRequester))

		response := versionDetailResolver.CreateVersionDetailDto(ctx)
		responseJson, _ := json.Marshal(response)

		mocks.CompareJson(t, responseJson, []byte(`{
			"version": "2.1.1",
			"endpoints": [{
				"identifier": "cdrs",
				"url": "http://localhost:8080/2.1.1/cdrs"
			}, {
				"identifier": "credentials",
				"url": "http://localhost:8080/2.1.1/credentials"
			}, {
				"identifier": "commands",
				"url": "http://localhost:8080/2.1.1/commands"
			}, {
				"identifier": "locations",
				"url": "http://localhost:8080/2.1.1/locations"
			}, {
				"identifier": "sessions",
				"url": "http://localhost:8080/2.1.1/sessions"
			}, {
				"identifier": "tariffs",
				"url": "http://localhost:8080/2.1.1/tariffs"
			}, {
				"identifier": "tokens",
				"url": "http://localhost:8080/2.1.1/tokens"
			}]
		}`))
	})
}
