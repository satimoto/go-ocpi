package mocks

import (
	"context"
	"errors"

	"github.com/satimoto/go-ocpi-api/ocpirpc"
	"google.golang.org/grpc"
)

func (s *MockOcpiService) UpdateTokens(ctx context.Context, in *ocpirpc.UpdateTokensRequest, opts ...grpc.CallOption) (*ocpirpc.UpdateTokensResponse, error) {
	if len(s.updateTokensMockData) == 0 {
		return &ocpirpc.UpdateTokensResponse{}, errors.New("NotFound")
	}

	response := s.updateTokensMockData[0]
	s.updateTokensMockData = s.updateTokensMockData[1:]
	return response, nil
}

func (s *MockOcpiService) SetUpdateTokensMockData(mockData *ocpirpc.UpdateTokensResponse) {
	s.updateTokensMockData = append(s.updateTokensMockData, mockData)
}