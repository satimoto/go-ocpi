package mocks

import (
	"context"
	"errors"

	"github.com/satimoto/go-ocpi-api/ocpirpc"
	"google.golang.org/grpc"
)

func (s *MockOcpiService) CdrCreated(ctx context.Context, in *ocpirpc.CdrCreatedRequest, opts ...grpc.CallOption) (*ocpirpc.CdrCreatedResponse, error) {
	if len(s.cdrCreatedMockData) == 0 {
		return &ocpirpc.CdrCreatedResponse{}, errors.New("NotFound")
	}

	response := s.cdrCreatedMockData[0]
	s.cdrCreatedMockData = s.cdrCreatedMockData[1:]
	return response, nil
}

func (s *MockOcpiService) SetCdrCreatedMockData(mockData *ocpirpc.CdrCreatedResponse) {
	s.cdrCreatedMockData = append(s.cdrCreatedMockData, mockData)
}
