package mocks

import (
	"context"
	"errors"

	"github.com/satimoto/go-ocpi/ocpirpc"
	"google.golang.org/grpc"
)

func (s *MockOcpiService) TestConnection(ctx context.Context, in *ocpirpc.TestConnectionRequest, opts ...grpc.CallOption) (*ocpirpc.TestConnectionResponse, error) {
	if len(s.testConnectionMockData) == 0 {
		return &ocpirpc.TestConnectionResponse{}, errors.New("NotFound")
	}

	response := s.testConnectionMockData[0]
	s.testConnectionMockData = s.testConnectionMockData[1:]
	return response, nil
}

func (s *MockOcpiService) TestMessage(ctx context.Context, in *ocpirpc.TestMessageRequest, opts ...grpc.CallOption) (*ocpirpc.TestMessageResponse, error) {
	if len(s.testMessageMockData) == 0 {
		return &ocpirpc.TestMessageResponse{}, errors.New("NotFound")
	}

	response := s.testMessageMockData[0]
	s.testMessageMockData = s.testMessageMockData[1:]
	return response, nil
}

func (s *MockOcpiService) SetTestConnectionMockData(mockData *ocpirpc.TestConnectionResponse) {
	s.testConnectionMockData = append(s.testConnectionMockData, mockData)
}

func (s *MockOcpiService) SetTestMessageMockData(mockData *ocpirpc.TestMessageResponse) {
	s.testMessageMockData = append(s.testMessageMockData, mockData)
}
