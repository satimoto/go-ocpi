package ocpi

import (
	"context"

	"github.com/satimoto/go-ocpi/ocpirpc"
	"google.golang.org/grpc"
)

func (s *OcpiService) CreateCredential(ctx context.Context, in *ocpirpc.CreateCredentialRequest, opts ...grpc.CallOption) (*ocpirpc.CreateCredentialResponse, error) {
	return s.getCredentialClient().CreateCredential(ctx, in, opts...)
}

func (s *OcpiService) RegisterCredential(ctx context.Context, in *ocpirpc.RegisterCredentialRequest, opts ...grpc.CallOption) (*ocpirpc.RegisterCredentialResponse, error) {
	return s.getCredentialClient().RegisterCredential(ctx, in, opts...)
}

func (s *OcpiService) UnregisterCredential(ctx context.Context, in *ocpirpc.UnregisterCredentialRequest, opts ...grpc.CallOption) (*ocpirpc.UnregisterCredentialResponse, error) {
	return s.getCredentialClient().UnregisterCredential(ctx, in, opts...)
}

func (s *OcpiService) getCredentialClient() ocpirpc.CredentialServiceClient {
	if s.credentialClient == nil {
		client := ocpirpc.NewCredentialServiceClient(s.clientConn)
		s.credentialClient = &client
	}

	return *s.credentialClient
}
