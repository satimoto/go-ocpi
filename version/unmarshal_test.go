package version_test

import (
	"context"
	"encoding/json"
	"io"
	"os"
	"strings"
	"testing"

	dbMocks "github.com/satimoto/go-datastore-mocks/db"
	"github.com/satimoto/go-ocpi-api/mocks"
	"github.com/satimoto/go-ocpi-api/rest"
	utilMocks "github.com/satimoto/go-ocpi-api/util/mocks"
	versionMocks "github.com/satimoto/go-ocpi-api/version/mocks"
)

func TestUnmarshalResponse(t *testing.T) {
	ctx := context.Background()

	os.Setenv("API_DOMAIN", "http://localhost:8080")
	defer os.Unsetenv("API_DOMAIN")

	t.Run("Success response", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &utilMocks.MockHTTPRequester{}
		versionResolver := versionMocks.NewResolver(mockRepository, utilMocks.NewOCPIRequester(mockHTTPRequester))

		payload := versionResolver.CreateVersionListPayload(ctx)
		response := rest.OCPISuccess(payload)
		responseJson, _ := json.Marshal(response)
		readerCloser := io.NopCloser(strings.NewReader(string(responseJson)))
		unmarshalReponse, _ := versionResolver.UnmarshalResponse(readerCloser)
		versionPayload := unmarshalReponse.Data
		versionJson, _ := json.Marshal(versionPayload)

		mocks.CompareJson(t, versionJson, []byte(`[{
			"version": "2.1.1",
			"url": "http://localhost:8080/2.1.1"
		}]`))
	})

	t.Run("Error response", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &utilMocks.MockHTTPRequester{}
		versionResolver := versionMocks.NewResolver(mockRepository, utilMocks.NewOCPIRequester(mockHTTPRequester))

		response := rest.OCPIRegistrationError(nil)
		responseJson, _ := json.Marshal(response)
		readerCloser := io.NopCloser(strings.NewReader(string(responseJson)))
		unmarshalReponse, _ := versionResolver.UnmarshalResponse(readerCloser)

		if unmarshalReponse.StatusCode != rest.STATUS_CODE_REGISTRATION_ERROR {
			t.Errorf("StatusCode mismatch: %v", unmarshalReponse.StatusCode)
		}
	})
}
