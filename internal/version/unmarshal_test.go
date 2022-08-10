package version_test

import (
	"context"
	"encoding/json"
	"io"
	"os"
	"strings"
	"testing"

	dbMocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	"github.com/satimoto/go-ocpi/internal/transportation"
	transportationMocks "github.com/satimoto/go-ocpi/internal/transportation/mocks"
	versionMocks "github.com/satimoto/go-ocpi/internal/version/mocks"
	"github.com/satimoto/go-ocpi/test/mocks"
)

func TestUnmarshalResponse(t *testing.T) {
	ctx := context.Background()

	os.Setenv("API_DOMAIN", "http://localhost:8080")
	defer os.Unsetenv("API_DOMAIN")

	t.Run("Success response", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		versionResolver := versionMocks.NewResolverWithServices(mockRepository, transportationMocks.NewOcpiRequester(mockHTTPRequester))

		dto := versionResolver.CreateVersionListDto(ctx)
		response := transportation.OcpiSuccess(dto)
		responseJson, _ := json.Marshal(response)
		readerCloser := io.NopCloser(strings.NewReader(string(responseJson)))
		unmarshalReponse, _ := versionResolver.UnmarshalPullDto(readerCloser)
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
		versionResolver := versionMocks.NewResolverWithServices(mockRepository, transportationMocks.NewOcpiRequester(mockHTTPRequester))

		response := transportation.OcpiRegistrationError(nil)
		responseJson, _ := json.Marshal(response)
		readerCloser := io.NopCloser(strings.NewReader(string(responseJson)))
		unmarshalReponse, _ := versionResolver.UnmarshalPullDto(readerCloser)

		if unmarshalReponse.StatusCode != transportation.STATUS_CODE_REGISTRATION_ERROR {
			t.Errorf("StatusCode mismatch: %v", unmarshalReponse.StatusCode)
		}
	})
}
