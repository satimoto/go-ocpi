package mocks

import (
	"github.com/satimoto/go-ocpi-api/ocpirpc"
)

type MockOcpiService struct {
	stopSessionMockData  []*ocpirpc.StopSessionResponse
	updateTokensMockData []*ocpirpc.UpdateTokensResponse
}

func NewService() *MockOcpiService {
	return &MockOcpiService{}
}
