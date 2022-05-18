package mocks

import (
	"github.com/satimoto/go-ocpi-api/ocpirpc"
)

type MockOcpiService struct {
	cdrCreatedMockData           []*ocpirpc.CdrCreatedResponse
	createCredentialMockData     []*ocpirpc.CreateCredentialResponse
	registerCredentialMockData   []*ocpirpc.RegisterCredentialResponse
	unregisterCredentialMockData []*ocpirpc.UnregisterCredentialResponse
	startSessionMockData         []*ocpirpc.StartSessionResponse
	stopSessionMockData          []*ocpirpc.StopSessionResponse
	sessionCreatedMockData       []*ocpirpc.SessionCreatedResponse
	createTokenMockData          []*ocpirpc.CreateTokenResponse
	updateTokensMockData         []*ocpirpc.UpdateTokensResponse
}

func NewService() *MockOcpiService {
	return &MockOcpiService{}
}
