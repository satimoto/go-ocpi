package ocpi

import (
	"context"

	"github.com/satimoto/go-ocpi-api/ocpirpc"
	"google.golang.org/grpc"
)

func (s *OcpiService) SessionCreated(ctx context.Context, in *ocpirpc.SessionCreatedRequest, opts ...grpc.CallOption) (*ocpirpc.SessionCreatedResponse, error) {
	return s.getSessionClient().SessionCreated(ctx, in, opts...)
}

func (s *OcpiService) getSessionClient() ocpirpc.SessionServiceClient {
	if s.sessionClient == nil {
		client := ocpirpc.NewSessionServiceClient(s.clientConn)
		s.sessionClient = &client
	}

	return *s.sessionClient
}
