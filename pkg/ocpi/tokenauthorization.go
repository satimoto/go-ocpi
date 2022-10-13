package ocpi

import (
	"context"

	"github.com/satimoto/go-ocpi/ocpirpc"
	"google.golang.org/grpc"
)

func (s *OcpiService) TokenAuthorizationCreated(ctx context.Context, in *ocpirpc.TokenAuthorizationCreatedRequest, opts ...grpc.CallOption) (*ocpirpc.TokenAuthorizationCreatedResponse, error) {
	return s.getTokenAuthorizationClient().TokenAuthorizationCreated(ctx, in, opts...)
}

func (s *OcpiService) getTokenAuthorizationClient() ocpirpc.TokenAuthorizationServiceClient {
	if s.tokenClient == nil {
		client := ocpirpc.NewTokenAuthorizationServiceClient(s.clientConn)
		s.tokenAuthorizationClient = &client
	}

	return *s.tokenAuthorizationClient
}
