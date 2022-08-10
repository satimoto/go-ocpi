package token_test

import (
	"context"
	"os"
	"testing"

	dbMocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	tokenMocks "github.com/satimoto/go-ocpi/internal/token/v2.1.1/mocks"
	transportationMocks "github.com/satimoto/go-ocpi/internal/transportation/mocks"
	"github.com/satimoto/go-ocpi/test/mocks"
)

func TestTokenProcess(t *testing.T) {
	os.Setenv("COUNTRY_CODE", "DE")
	os.Setenv("PARTY_ID", "SAT")
	defer os.Unsetenv("COUNTRY_CODE")
	defer os.Unsetenv("PARTY_ID")

	ctx := context.Background()

	t.Run("Generate AuthID", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		tokenResolver := tokenMocks.NewResolverWithServices(mockRepository, transportationMocks.NewOcpiRequester(mockHTTPRequester))

		for i := 0; i < 1000; i++ {
			authID, err := tokenResolver.GenerateAuthID(ctx)
			t.Logf("AuthId: %v", authID)

			if err != nil {
				t.Errorf("Error: %v %v", authID, err)
			}
		}
	})
}