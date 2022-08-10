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

func (s *MockOcpiService) SessionUpdated(ctx context.Context, in *ocpirpc.SessionUpdatedRequest, opts ...grpc.CallOption) (*ocpirpc.SessionUpdatedResponse, error) {
	if len(s.sessionUpdatedMockData) == 0 {
		return &ocpirpc.SessionUpdatedResponse{}, errors.New("NotFound")
	}

	response := s.sessionUpdatedMockData[0]
	s.sessionUpdatedMockData = s.sessionUpdatedMockData[1:]
	return response, nil
}

func (s *MockOcpiService) SetSessionCreatedMockData(mockData *ocpirpc.SessionCreatedResponse) {
	s.sessionCreatedMockData = append(s.sessionCreatedMockData, mockData)
}

func (s *MockOcpiService) SetSessionUpdatedMockData(mockData *ocpirpc.SessionUpdatedResponse) {
	s.sessionUpdatedMockData = append(s.sessionUpdatedMockData, mockData)
}