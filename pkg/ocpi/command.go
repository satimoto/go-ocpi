package ocpi

import (
	"context"

	"github.com/satimoto/go-ocpi-api/ocpirpc"
	"google.golang.org/grpc"
)

func (s *OcpiService) StartSession(ctx context.Context, in *ocpirpc.StartSessionRequest, opts ...grpc.CallOption) (*ocpirpc.StartSessionResponse, error) {
	return s.getCommandClient().StartSession(ctx, in, opts...)
}

func (s *OcpiService) StopSession(ctx context.Context, in *ocpirpc.StopSessionRequest, opts ...grpc.CallOption) (*ocpirpc.StopSessionResponse, error) {
	return s.getCommandClient().StopSession(ctx, in, opts...)
}

func (s *OcpiService) getCommandClient() ocpirpc.CommandServiceClient {
	if s.commandClient == nil {
		client := ocpirpc.NewCommandServiceClient(s.clientConn)
		s.commandClient = &client
	}

	return *s.commandClient
}
