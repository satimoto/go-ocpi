package ocpi

import (
	"context"

	"github.com/satimoto/go-ocpi/ocpirpc"
	"google.golang.org/grpc"
)

func (s *OcpiService) CdrCreated(ctx context.Context, in *ocpirpc.CdrCreatedRequest, opts ...grpc.CallOption) (*ocpirpc.CdrCreatedResponse, error) {
	return s.getCdrClient().CdrCreated(ctx, in, opts...)
}

func (s *OcpiService) getCdrClient() ocpirpc.CdrServiceClient {
	if s.cdrClient == nil {
		client := ocpirpc.NewCdrServiceClient(s.clientConn)
		s.cdrClient = &client
	}

	return *s.cdrClient
}
