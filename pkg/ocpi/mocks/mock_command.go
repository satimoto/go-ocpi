package mocks

import (
	"context"
	"errors"

	"github.com/satimoto/go-ocpi-api/ocpirpc"
	"google.golang.org/grpc"
)

func (s *MockOcpiService) StopSession(ctx context.Context, in *ocpirpc.StopSessionRequest, opts ...grpc.CallOption) (*ocpirpc.StopSessionResponse, error) {
	if len(s.stopSessionMockData) == 0 {
		return &ocpirpc.StopSessionResponse{}, errors.New("NotFound")
	}

	response := s.stopSessionMockData[0]
	s.stopSessionMockData = s.stopSessionMockData[1:]
	return response, nil
}

func (s *MockOcpiService) SetStopSessionMockData(mockData *ocpirpc.StopSessionResponse) {
	s.stopSessionMockData = append(s.stopSessionMockData, mockData)
}
