package ocpi

import (
	"context"
	"log"
	"time"

	"github.com/satimoto/go-ocpi/ocpirpc"
	"google.golang.org/grpc"
)

func (s *OcpiService) TestConnection(ctx context.Context, in *ocpirpc.TestConnectionRequest, opts ...grpc.CallOption) (*ocpirpc.TestConnectionResponse, error) {
	timerStart := time.Now()
	response, err := s.getRpcClient().TestConnection(ctx, in, opts...)
	timerStop := time.Now()

	log.Printf("TestConnection responded in %f seconds", timerStop.Sub(timerStart).Seconds())

	return response, err
}

func (s *OcpiService) TestMessage(ctx context.Context, in *ocpirpc.TestMessageRequest, opts ...grpc.CallOption) (*ocpirpc.TestMessageResponse, error) {
	timerStart := time.Now()
	response, err := s.getRpcClient().TestMessage(ctx, in, opts...)
	timerStop := time.Now()

	log.Printf("TestMessage responded in %f seconds", timerStop.Sub(timerStart).Seconds())

	return response, err
}

func (s *OcpiService) getRpcClient() ocpirpc.RpcServiceClient {
	if s.rpcClient == nil {
		client := ocpirpc.NewRpcServiceClient(s.clientConn)
		s.rpcClient = &client
	}

	return *s.rpcClient
}
