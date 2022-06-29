package mocks

import (
	"context"
	"errors"

	"github.com/satimoto/go-ocpi/ocpirpc"
	"google.golang.org/grpc"
)

func (s *MockOcpiService) SessionCreated(ctx context.Context, in *ocpirpc.SessionCreatedRequest, opts ...grpc.CallOption) (*ocpirpc.SessionCreatedResponse, error) {
	if len(s.sessionCreatedMockData) == 0 {
		return &ocpirpc.SessionCreatedResponse{}, errors.New("NotFound")
	}

	response := s.sessionCreatedMockData[0]
	s.sessionCreatedMockData = s.sessionCreatedMockData[1:]
	return response, nil
}

func (s *MockOcpiService) SetSessionCreatedMockData(mockData *ocpirpc.SessionCreatedResponse) {
	s.sessionCreatedMockData = append(s.sessionCreatedMockData, mockData)
}
