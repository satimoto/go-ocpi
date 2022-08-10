package ocpi

import (
	"context"

	"github.com/satimoto/go-ocpi/ocpirpc"
	"google.golang.org/grpc"
)

func (s *OcpiService) TestConnection(ctx context.Context, in *ocpirpc.TestConnectionRequest, opts ...grpc.CallOption) (*ocpirpc.TestConnectionResponse, error) {
	return s.getRpcClient().TestConnection(ctx, in, opts...)
}

func (s *OcpiService) TestMessage(ctx context.Context, in *ocpirpc.TestMessageRequest, opts ...grpc.CallOption) (*ocpirpc.TestMessageResponse, error) {
	return s.getRpcClient().TestMessage(ctx, in, opts...)
}

func (s *OcpiService) getRpcClient() ocpirpc.RpcServiceClient {
	if s.rpcClient == nil {
		client := ocpirpc.NewRpcServiceClient(s.clientConn)
		s.rpcClient = &client
	}

	return *s.rpcClient
}
