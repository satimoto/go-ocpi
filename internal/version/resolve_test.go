package version_test

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"

	dbMocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	"github.com/satimoto/go-datastore/pkg/util"
	"github.com/satimoto/go-ocpi/internal/transportation"
	transportationMocks "github.com/satimoto/go-ocpi/internal/transportation/mocks"
	versionMocks "github.com/satimoto/go-ocpi/internal/version/mocks"
	"github.com/satimoto/go-ocpi/test/mocks"
)

func TestPullVersions(t *testing.T) {
	ctx := context.Background()

	os.Setenv("API_DOMAIN", "http://localhost:8080")
	defer os.Unsetenv("API_DOMAIN")

	t.Run("Success request", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		versionResolver := versionMocks.NewResolver(mockRepository, transportationMocks.NewOcpiService(mockHTTPRequester))

		bodyBytes := `{
			"data": [{
				"version": "2.1.1",
				"url": "http://localhost:8080/2.1.1"
			}],
			"status_code": 1000,
			"status_message": "Success",
			"timestamp": "2018-12-16T11:00:02Z"
		}`
		readerCloser := io.NopCloser(strings.NewReader(string(bodyBytes)))

		mockHTTPRequester.SetResponse(mocks.MockResponseData{
			Response: &http.Response{
				StatusCode: 200,
				Body:       readerCloser,
			},
		})

		header := transportation.OcpiRequestHeader{
			Authorization: util.NilString("F72FB7A3-BD45-4A9E-8972-D0452EA0F452"),
			ToCountryCode: util.NilString("DE"),
			ToPartyId:     util.NilString("EXA"),
		}

		response := versionResolver.PullVersions(ctx, "/versions", header, 1)
		responseJson, _ := json.Marshal(response)

		mocks.CompareJson(t, responseJson, []byte(`[{
			"id": 0,
			"credentialID": 1,
			"version": "2.1.1",
			"url": "http://localhost:8080/2.1.1"
		}]`))
	})
}
