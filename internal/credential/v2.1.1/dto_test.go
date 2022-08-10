package credential_test

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	"github.com/satimoto/go-datastore/pkg/db"
	dbMocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	"github.com/satimoto/go-datastore/pkg/util"
	credentialMocks "github.com/satimoto/go-ocpi/internal/credential/v2.1.1/mocks"
	transportationMocks "github.com/satimoto/go-ocpi/internal/transportation/mocks"
	"github.com/satimoto/go-ocpi/test/mocks"
)

func TestCreateCredentialDto(t *testing.T) {
	ctx := context.Background()

	os.Setenv("API_DOMAIN", "https://api.local:8080")
	os.Setenv("PARTY_ID", "SAT")
	os.Setenv("COUNTRY_CODE", "DE")
	os.Setenv("WEB_DOMAIN", "https://web.local:8080")
	defer os.Unsetenv("API_DOMAIN")
	defer os.Unsetenv("PARTY_ID")
	defer os.Unsetenv("COUNTRY_CODE")
	defer os.Unsetenv("WEB_DOMAIN")

	t.Run("Create dto", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		credentialResolver := credentialMocks.NewResolverWithServices(mockRepository, transportationMocks.NewOcpiRequester(mockHTTPRequester))

		credential := db.Credential{
			ClientToken:      util.SqlNullString("EF3ABC19-84AB-476D-A12D-17FA42FB3CE5"),
			ServerToken:      util.SqlNullString("001DEC16-EAA1-4B7C-9884-AC45F248E4D7"),
			PartyID:          "BUS",
			CountryCode:      "DE",
			BusinessDetailID: 123,
			LastUpdated:      *util.ParseTime("2015-06-29T20:39:09Z", nil),
		}

		response := credentialResolver.CreateCredentialDto(ctx, credential)
		responseJson, _ := json.Marshal(response)

		mocks.CompareJson(t, responseJson, []byte(`{
			"token": "001DEC16-EAA1-4B7C-9884-AC45F248E4D7",
			"url": "https://api.local:8080",
			"party_id": "SAT",
			"country_code": "DE",
			"business_details": {
				"name": "Satimoto",
				"website": "https://web.local:8080",
				"logo": {
					"url": "https://web.local:8080/logo.png",
					"thumbnail": "https://web.local:8080/logo-thumb.png",
					"category": "OPERATOR",
					"type": "png",
					"width": 512,
					"height": 512
				}
			}
		}`))
	})
}