package mocks

import (
	"context"
	"errors"

	"github.com/satimoto/go-ocpi/ocpirpc"
	"google.golang.org/grpc"
)

func (s *MockOcpiService) CreateToken(ctx context.Context, in *ocpirpc.CreateTokenRequest, opts ...grpc.CallOption) (*ocpirpc.CreateTokenResponse, error) {
	if len(s.createTokenMockData) == 0 {
		return &ocpirpc.CreateTokenResponse{}, errors.New("NotFound")
	}

	response := s.createTokenMockData[0]
	s.createTokenMockData = s.createTokenMockData[1:]
	return response, nil
}

func (s *MockOcpiService) UpdateTokens(ctx context.Context, in *ocpirpc.UpdateTokensRequest, opts ...grpc.CallOption) (*ocpirpc.UpdateTokensResponse, error) {
	if len(s.updateTokensMockData) == 0 {
		return &ocpirpc.UpdateTokensResponse{}, errors.New("NotFound")
	}

	response := s.updateTokensMockData[0]
	s.updateTokensMockData = s.updateTokensMockData[1:]
	return response, nil
}

func (s *MockOcpiService) SetCreateTokenMockData(mockData *ocpirpc.CreateTokenResponse) {
	s.createTokenMockData = append(s.createTokenMockData, mockData)
}

func (s *MockOcpiService) SetUpdateTokensMockData(mockData *ocpirpc.UpdateTokensResponse) {
	s.updateTokensMockData = append(s.updateTokensMockData, mockData)
}
