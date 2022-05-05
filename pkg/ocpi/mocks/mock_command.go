package mocks

import (
	"context"
	"errors"

	"github.com/satimoto/go-ocpi-api/ocpirpc"
	"google.golang.org/grpc"
)

func (s *MockOcpiService) StartSession(ctx context.Context, in *ocpirpc.StartSessionRequest, opts ...grpc.CallOption) (*ocpirpc.StartSessionResponse, error) {
	if len(s.startSessionMockData) == 0 {
		return &ocpirpc.StartSessionResponse{}, errors.New("NotFound")
	}

	response := s.startSessionMockData[0]
	s.startSessionMockData = s.startSessionMockData[1:]
	return response, nil
}

func (s *MockOcpiService) StopSession(ctx context.Context, in *ocpirpc.StopSessionRequest, opts ...grpc.CallOption) (*ocpirpc.StopSessionResponse, error) {
	if len(s.stopSessionMockData) == 0 {
		return &ocpirpc.StopSessionResponse{}, errors.New("NotFound")
	}

	response := s.stopSessionMockData[0]
	s.stopSessionMockData = s.stopSessionMockData[1:]
	return response, nil
}

func (s *MockOcpiService) SetStartSessionMockData(mockData *ocpirpc.StartSessionResponse) {
	s.startSessionMockData = append(s.startSessionMockData, mockData)
}

func (s *MockOcpiService) SetStopSessionMockData(mockData *ocpirpc.StopSessionResponse) {
	s.stopSessionMockData = append(s.stopSessionMockData, mockData)
}
