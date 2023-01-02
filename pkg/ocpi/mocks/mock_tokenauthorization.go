package mocks

import (
	"context"
	"errors"

	"github.com/satimoto/go-ocpi/ocpirpc"
	"google.golang.org/grpc"
)

func (s *MockOcpiService) UpdateTokenAuthorization(ctx context.Context, in *ocpirpc.UpdateTokenAuthorizationRequest, opts ...grpc.CallOption) (*ocpirpc.UpdateTokenAuthorizationResponse, error) {
	if len(s.updateTokenAuthorizationMockData) == 0 {
		return &ocpirpc.UpdateTokenAuthorizationResponse{}, errors.New("NotFound")
	}

	response := s.updateTokenAuthorizationMockData[0]
	s.updateTokenAuthorizationMockData = s.updateTokenAuthorizationMockData[1:]
	return response, nil
}

func (s *MockOcpiService) SetTokenAuthorizationCreatedMockData(mockData *ocpirpc.UpdateTokenAuthorizationResponse) {
	s.updateTokenAuthorizationMockData = append(s.updateTokenAuthorizationMockData, mockData)
}
