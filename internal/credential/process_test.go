package credential_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/satimoto/go-datastore/pkg/db"
	dbMocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	"github.com/satimoto/go-datastore/pkg/util"
	credentialMocks "github.com/satimoto/go-ocpi-api/internal/credential/mocks"
	transportationMocks "github.com/satimoto/go-ocpi-api/internal/transportation/mocks"
	"github.com/satimoto/go-ocpi-api/test/mocks"
)

func TestRegisterCredential(t *testing.T) {
	ctx := context.Background()

	t.Run("Empty token", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		credentialResolver := credentialMocks.NewResolverWithServices(mockRepository, transportationMocks.NewOcpiRequester(mockHTTPRequester))

		credential := db.Credential{
			ClientToken: util.SqlNullString(nil),
			Url:         "http://localhost:9000/versions",
			CountryCode: "DE",
			PartyID:     "ABC",
		}

		_, err := credentialResolver.RegisterCredential(ctx, credential, credential.ClientToken.String)

		if err == nil || err.Error() != "Registration error" {
			t.Errorf("Error mismatch: '%v' expecting '%v'", err, "Registration error")
		}
	})

	t.Run("No versions", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		credentialResolver := credentialMocks.NewResolverWithServices(mockRepository, transportationMocks.NewOcpiRequester(mockHTTPRequester))

		credential := db.Credential{
			ClientToken: util.SqlNullString("1802EC4A-2A34-4573-803E-1E142CF7BC1C"),
			Url:         "http://localhost:9000/versions",
			CountryCode: "DE",
			PartyID:     "ABC",
		}

		_, err := credentialResolver.RegisterCredential(ctx, credential, credential.ClientToken.String)

		if err == nil || err.Error() != "Unsupported version" {
			t.Errorf("Error mismatch: '%v' expecting '%v'", err, "Unsupported version")
		}
	})

	t.Run("Unsupported version", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		credentialResolver := credentialMocks.NewResolverWithServices(mockRepository, transportationMocks.NewOcpiRequester(mockHTTPRequester))

		token := "1802EC4A-2A34-4573-803E-1E142CF7BC1C"
		credential := db.Credential{
			ClientToken: util.SqlNullString(nil),
			Url:         "http://localhost:9000/versions",
			CountryCode: "DE",
			PartyID:     "ABC",
		}

		versions := []db.Version{}
		versions = append(versions, db.Version{
			Version: "2.0",
			Url:     "http://localhost:9000/2.0",
		})

		mockRepository.SetListVersionsMockData(dbMocks.VersionsMockData{Versions: versions})

		_, err := credentialResolver.RegisterCredential(ctx, credential, token)

		if err == nil || err.Error() != "Unsupported version" {
			t.Errorf("Error mismatch: '%v' expecting '%v'", err, "Unsupported version")
		}
	})

	t.Run("Unsupported version", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		credentialResolver := credentialMocks.NewResolverWithServices(mockRepository, transportationMocks.NewOcpiRequester(mockHTTPRequester))

		token := "1802EC4A-2A34-4573-803E-1E142CF7BC1C"
		credential := db.Credential{
			ClientToken: util.SqlNullString(nil),
			ServerToken: util.SqlNullString("C33B6CB9-BDE4-4C47-91A2-DC390DA3C374"),
			Url:         "http://localhost:9000/versions",
			CountryCode: "DE",
			PartyID:     "ABC",
		}

		wrongVersion := db.Version{
			Version: "2.0",
			Url:     "http://localhost:9000/2.0",
		}
		rightVersion := db.Version{
			Version: "2.1.1",
			Url:     "http://localhost:9000/2.1.1",
		}
		versions := []db.Version{wrongVersion, rightVersion}

		mockRepository.SetListVersionsMockData(dbMocks.VersionsMockData{Versions: versions})

		_, err := credentialResolver.RegisterCredential(ctx, credential, token)

		if err != nil {
			t.Errorf("Error mismatch: '%v' expecting '%v'", err, nil)
		}

		mockHTTPRequester.GetRequest()            // Versions
		request := mockHTTPRequester.GetRequest() // Version endpoints

		authenticationHeader := request.Header.Get("Authentication")
		expectedAuthenticationHeader := fmt.Sprintf("Token %s", token)

		if authenticationHeader != expectedAuthenticationHeader {
			t.Errorf("Error mismatch: '%v' expecting '%v'", authenticationHeader, expectedAuthenticationHeader)
		}

		if request.URL.String() != rightVersion.Url {
			t.Errorf("Error mismatch: '%v' expecting '%v'", request.URL.String(), rightVersion.Url)
		}

		updateCredential, err := mockRepository.GetUpdateCredentialMockData()

		if updateCredential.ClientToken.String != token {
			t.Errorf("Error mismatch: '%v' expecting '%v'", updateCredential.ClientToken.String, token)
		}

		if updateCredential.ServerToken.String == credential.ServerToken.String {
			t.Errorf("Error mismatch: '%v' expecting new UUID", updateCredential.ServerToken.String)
		}
	})
}

func TestUnregisterCredential(t *testing.T) {
	ctx := context.Background()

	t.Run("Empty token", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		credentialResolver := credentialMocks.NewResolverWithServices(mockRepository, transportationMocks.NewOcpiRequester(mockHTTPRequester))

		credential := db.Credential{
			ServerToken: util.SqlNullString("D501D324-A33A-41E0-91DF-34A73BB8F8A7"),
			Url:         "http://localhost:9000/versions",
			CountryCode: "DE",
			PartyID:     "ABC",
		}

		_, err := credentialResolver.UnregisterCredential(ctx, credential)

		if err == nil || err.Error() != "error credential not registered" {
			t.Errorf("Error mismatch: '%v' expecting '%v'", err, "Error credential not registered")
		}
	})

	t.Run("No version endpoint", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		credentialResolver := credentialMocks.NewResolverWithServices(mockRepository, transportationMocks.NewOcpiRequester(mockHTTPRequester))

		credential := db.Credential{
			ClientToken: util.SqlNullString("1802EC4A-2A34-4573-803E-1E142CF7BC1C"),
			ServerToken: util.SqlNullString("D501D324-A33A-41E0-91DF-34A73BB8F8A7"),
			Url:         "http://localhost:9000/versions",
			CountryCode: "DE",
			PartyID:     "ABC",
		}

		_, err := credentialResolver.UnregisterCredential(ctx, credential)

		if err == nil || err.Error() != "error retrieving version endpoint" {
			t.Errorf("Error mismatch: '%v' expecting '%v'", err, "Error retrieving version endpoint")
		}
	})

	t.Run("Bad response", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		credentialResolver := credentialMocks.NewResolverWithServices(mockRepository, transportationMocks.NewOcpiRequester(mockHTTPRequester))

		credential := db.Credential{
			ClientToken: util.SqlNullString("1802EC4A-2A34-4573-803E-1E142CF7BC1C"),
			ServerToken: util.SqlNullString("D501D324-A33A-41E0-91DF-34A73BB8F8A7"),
			Url:         "http://localhost:9000/versions",
			CountryCode: "DE",
			PartyID:     "ABC",
		}

		mockRepository.SetGetVersionEndpointByIdentityMockData(dbMocks.VersionEndpointMockData{VersionEndpoint: db.VersionEndpoint{
			Identifier: "credentials",
			Url:        "http://localhost:9000/2.1.1/credentials",
		}})

		mockHTTPRequester.SetResponseWithBytes(200, ``, nil)

		_, err := credentialResolver.UnregisterCredential(ctx, credential)

		if err == nil || err.Error() != "error unmarshalling response" {
			t.Errorf("Error mismatch: '%v' expecting '%v'", err, "Error unmarshalling response")
		}
	})

	t.Run("Unsuccessful response", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		credentialResolver := credentialMocks.NewResolverWithServices(mockRepository, transportationMocks.NewOcpiRequester(mockHTTPRequester))

		credential := db.Credential{
			ClientToken: util.SqlNullString("1802EC4A-2A34-4573-803E-1E142CF7BC1C"),
			ServerToken: util.SqlNullString("D501D324-A33A-41E0-91DF-34A73BB8F8A7"),
			Url:         "http://localhost:9000/versions",
			CountryCode: "DE",
			PartyID:     "ABC",
		}

		mockRepository.SetGetVersionEndpointByIdentityMockData(dbMocks.VersionEndpointMockData{VersionEndpoint: db.VersionEndpoint{
			Identifier: "credentials",
			Url:        "http://localhost:9000/2.1.1/credentials",
		}})

		mockHTTPRequester.SetResponseWithBytes(200, `{
			"status_code": 2003,
			"status_message": "Unknown resource",
			"timestamp": "2018-12-16T11:00:02Z"
		}`, nil)

		_, err := credentialResolver.UnregisterCredential(ctx, credential)

		if err == nil || err.Error() != "error in delete response" {
			t.Errorf("Error mismatch: '%v' expecting '%v'", err, "Error in delete response")
		}
	})

	t.Run("Successful response", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		credentialResolver := credentialMocks.NewResolverWithServices(mockRepository, transportationMocks.NewOcpiRequester(mockHTTPRequester))

		credential := db.Credential{
			ClientToken: util.SqlNullString("1802EC4A-2A34-4573-803E-1E142CF7BC1C"),
			ServerToken: util.SqlNullString("D501D324-A33A-41E0-91DF-34A73BB8F8A7"),
			Url:         "http://localhost:9000/versions",
			CountryCode: "DE",
			PartyID:     "ABC",
		}

		mockRepository.SetGetVersionEndpointByIdentityMockData(dbMocks.VersionEndpointMockData{VersionEndpoint: db.VersionEndpoint{
			Identifier: "credentials",
			Url:        "http://localhost:9000/2.1.1/credentials",
		}})

		mockHTTPRequester.SetResponseWithBytes(200, `{
			"status_code": 1000,
			"status_message": "Success",
			"timestamp": "2018-12-16T11:00:02Z"
		}`, nil)

		_, err := credentialResolver.UnregisterCredential(ctx, credential)

		if err != nil {
			t.Errorf("Error mismatch: '%v' expecting '%v'", err, nil)
		}

		cred, _ := mockRepository.GetUpdateCredentialMockData()

		if cred.ServerToken.Valid || cred.ServerToken.String == "D501D324-A33A-41E0-91DF-34A73BB8F8A7" {
			t.Errorf("Error mismatch: '%v' expecting '%v'", cred.ClientToken.String, "")
		}
	})

}
