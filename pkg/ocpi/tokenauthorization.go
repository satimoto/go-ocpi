package ocpi

import (
	"context"

	"github.com/satimoto/go-ocpi/ocpirpc"
	"google.golang.org/grpc"
)

func (s *OcpiService) UpdateTokenAuthorization(ctx context.Context, in *ocpirpc.UpdateTokenAuthorizationRequest, opts ...grpc.CallOption) (*ocpirpc.UpdateTokenAuthorizationResponse, error) {
	return s.getTokenAuthorizationClient().UpdateTokenAuthorization(ctx, in, opts...)
}

func (s *OcpiService) getTokenAuthorizationClient() ocpirpc.TokenAuthorizationServiceClient {
	if s.tokenAuthorizationClient == nil {
		client := ocpirpc.NewTokenAuthorizationServiceClient(s.clientConn)
		s.tokenAuthorizationClient = &client
	}

	return *s.tokenAuthorizationClient
}
