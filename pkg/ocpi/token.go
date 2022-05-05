package ocpi

import (
	"context"

	"github.com/satimoto/go-ocpi-api/ocpirpc"
	"google.golang.org/grpc"
)

func (s *OcpiService) CreateToken(ctx context.Context, in *ocpirpc.CreateTokenRequest, opts ...grpc.CallOption) (*ocpirpc.CreateTokenResponse, error) {
	return s.getTokenClient().CreateToken(ctx, in, opts...)
}

func (s *OcpiService) UpdateTokens(ctx context.Context, in *ocpirpc.UpdateTokensRequest, opts ...grpc.CallOption) (*ocpirpc.UpdateTokensResponse, error) {
	return s.getTokenClient().UpdateTokens(ctx, in, opts...)
}

func (s *OcpiService) getTokenClient() ocpirpc.TokenServiceClient {
	if s.commandClient == nil {
		client := ocpirpc.NewTokenServiceClient(s.clientConn)
		s.tokenClient = &client
	}

	return *s.tokenClient
}
