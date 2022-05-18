package version_test

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	dbMocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	transportationMocks "github.com/satimoto/go-ocpi-api/internal/transportation/mocks"
	versionMocks "github.com/satimoto/go-ocpi-api/internal/version/mocks"
	"github.com/satimoto/go-ocpi-api/test/mocks"
)

func TestCreateDto(t *testing.T) {
	ctx := context.Background()

	os.Setenv("API_DOMAIN", "http://localhost:8080")
	defer os.Unsetenv("API_DOMAIN")

	t.Run("Create dto", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		versionResolver := versionMocks.NewResolverWithServices(mockRepository, transportationMocks.NewOcpiRequester(mockHTTPRequester))

		response := versionResolver.CreateVersionListDto(ctx)
		responseJson, _ := json.Marshal(response)

		mocks.CompareJson(t, responseJson, []byte(`[{
			"version": "2.1.1",
			"url": "http://localhost:8080/2.1.1"
		}]`))
	})
}
