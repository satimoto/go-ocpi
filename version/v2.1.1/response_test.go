package version_test

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	dbMocks "github.com/satimoto/go-datastore-mocks/db"
	"github.com/satimoto/go-ocpi-api/mocks"
	versionMocks "github.com/satimoto/go-ocpi-api/version/v2.1.1/mocks"
)

func TestCreateVersionDetailPayload(t *testing.T) {
	ctx := context.Background()

	os.Setenv("API_DOMAIN", "http://localhost:8080")
	defer os.Unsetenv("API_DOMAIN")

	t.Run("Create payload", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		versionResolver := versionMocks.NewResolver(mockRepository)

		response := versionResolver.CreateVersionDetailPayload(ctx)
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
