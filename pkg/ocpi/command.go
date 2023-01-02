package ocpi

import (
	"context"
	"log"
	"time"

	"github.com/satimoto/go-ocpi/ocpirpc"
	"google.golang.org/grpc"
)

func (s *OcpiService) StartSession(ctx context.Context, in *ocpirpc.StartSessionRequest, opts ...grpc.CallOption) (*ocpirpc.StartSessionResponse, error) {
	timerStart := time.Now()
	response, err := s.getCommandClient().StartSession(ctx, in, opts...)
	timerStop := time.Now()

	log.Printf("StartSession responded in %f seconds", timerStop.Sub(timerStart).Seconds())

	return response, err
}

func (s *OcpiService) StopSession(ctx context.Context, in *ocpirpc.StopSessionRequest, opts ...grpc.CallOption) (*ocpirpc.StopSessionResponse, error) {
	timerStart := time.Now()
	response, err := s.getCommandClient().StopSession(ctx, in, opts...)
	timerStop := time.Now()

	log.Printf("StopSession responded in %f seconds", timerStop.Sub(timerStart).Seconds())

	return response, err
}

func (s *OcpiService) getCommandClient() ocpirpc.CommandServiceClient {
	if s.commandClient == nil {
		client := ocpirpc.NewCommandServiceClient(s.clientConn)
		s.commandClient = &client
	}

	return *s.commandClient
}
