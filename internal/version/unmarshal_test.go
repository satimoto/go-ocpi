package version_test

import (
	"context"
	"encoding/json"
	"io"
	"os"
	"strings"
	"testing"

	dbMocks "github.com/satimoto/go-datastore-mocks/db"
	"github.com/satimoto/go-ocpi-api/internal/ocpi"
	versionMocks "github.com/satimoto/go-ocpi-api/internal/version/mocks"
	"github.com/satimoto/go-ocpi-api/test/mocks"
)

func TestUnmarshalResponse(t *testing.T) {
	ctx := context.Background()

	os.Setenv("API_DOMAIN", "http://localhost:8080")
	defer os.Unsetenv("API_DOMAIN")

	t.Run("Success response", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		versionResolver := versionMocks.NewResolver(mockRepository, mocks.NewOCPIRequester(mockHTTPRequester))

		dto := versionResolver.CreateVersionListDto(ctx)
		response := ocpi.OCPISuccess(dto)
		responseJson, _ := json.Marshal(response)
		readerCloser := io.NopCloser(strings.NewReader(string(responseJson)))
		unmarshalReponse, _ := versionResolver.UnmarshalResponse(readerCloser)
		versionDto := unmarshalReponse.Data
		versionJson, _ := json.Marshal(versionDto)

		mocks.CompareJson(t, versionJson, []byte(`[{
			"version": "2.1.1",
			"url": "http://localhost:8080/2.1.1"
		}]`))
	})

	t.Run("Error response", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		versionResolver := versionMocks.NewResolver(mockRepository, mocks.NewOCPIRequester(mockHTTPRequester))

		response := ocpi.OCPIRegistrationError(nil)
		responseJson, _ := json.Marshal(response)
		readerCloser := io.NopCloser(strings.NewReader(string(responseJson)))
		unmarshalReponse, _ := versionResolver.UnmarshalResponse(readerCloser)

		if unmarshalReponse.StatusCode != ocpi.STATUS_CODE_REGISTRATION_ERROR {
			t.Errorf("StatusCode mismatch: %v", unmarshalReponse.StatusCode)
		}
	})
}
