package mocks

import (
	"github.com/satimoto/go-ocpi/ocpirpc"
)

type MockOcpiService struct {
	cdrCreatedMockData               []*ocpirpc.CdrCreatedResponse
	createCredentialMockData         []*ocpirpc.CreateCredentialResponse
	registerCredentialMockData       []*ocpirpc.RegisterCredentialResponse
	syncCredentialMockData           []*ocpirpc.SyncCredentialResponse
	unregisterCredentialMockData     []*ocpirpc.UnregisterCredentialResponse
	startSessionMockData             []*ocpirpc.StartSessionResponse
	stopSessionMockData              []*ocpirpc.StopSessionResponse
	testConnectionMockData           []*ocpirpc.TestConnectionResponse
	testMessageMockData              []*ocpirpc.TestMessageResponse
	sessionCreatedMockData           []*ocpirpc.SessionCreatedResponse
	sessionUpdatedMockData           []*ocpirpc.SessionUpdatedResponse
	updateTokenAuthorizationMockData []*ocpirpc.UpdateTokenAuthorizationResponse
	createTokenMockData              []*ocpirpc.CreateTokenResponse
	updateTokensMockData             []*ocpirpc.UpdateTokensResponse
}

func NewService() *MockOcpiService {
	return &MockOcpiService{}
}
