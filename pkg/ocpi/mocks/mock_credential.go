package mocks

import (
	"context"
	"errors"

	"github.com/satimoto/go-ocpi/ocpirpc"
	"google.golang.org/grpc"
)

func (s *MockOcpiService) CreateCredential(ctx context.Context, in *ocpirpc.CreateCredentialRequest, opts ...grpc.CallOption) (*ocpirpc.CreateCredentialResponse, error) {
	if len(s.createCredentialMockData) == 0 {
		return &ocpirpc.CreateCredentialResponse{}, errors.New("NotFound")
	}

	response := s.createCredentialMockData[0]
	s.createCredentialMockData = s.createCredentialMockData[1:]
	return response, nil
}

func (s *MockOcpiService) RegisterCredential(ctx context.Context, in *ocpirpc.RegisterCredentialRequest, opts ...grpc.CallOption) (*ocpirpc.RegisterCredentialResponse, error) {
	if len(s.registerCredentialMockData) == 0 {
		return &ocpirpc.RegisterCredentialResponse{}, errors.New("NotFound")
	}

	response := s.registerCredentialMockData[0]
	s.registerCredentialMockData = s.registerCredentialMockData[1:]
	return response, nil
}

func (s *MockOcpiService) UnregisterCredential(ctx context.Context, in *ocpirpc.UnregisterCredentialRequest, opts ...grpc.CallOption) (*ocpirpc.UnregisterCredentialResponse, error) {
	if len(s.unregisterCredentialMockData) == 0 {
		return &ocpirpc.UnregisterCredentialResponse{}, errors.New("NotFound")
	}

	response := s.unregisterCredentialMockData[0]
	s.unregisterCredentialMockData = s.unregisterCredentialMockData[1:]
	return response, nil
}

func (s *MockOcpiService) SetCreateCredentialMockData(mockData *ocpirpc.CreateCredentialResponse) {
	s.createCredentialMockData = append(s.createCredentialMockData, mockData)
}

func (s *MockOcpiService) SetRegisterCredentialMockData(mockData *ocpirpc.RegisterCredentialResponse) {
	s.registerCredentialMockData = append(s.registerCredentialMockData, mockData)
}

func (s *MockOcpiService) SetUnregisterCredentialMockData(mockData *ocpirpc.UnregisterCredentialResponse) {
	s.unregisterCredentialMockData = append(s.unregisterCredentialMockData, mockData)
}
