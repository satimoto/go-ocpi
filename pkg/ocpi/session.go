package ocpi

import (
	"context"
	"log"
	"time"

	"github.com/satimoto/go-ocpi/ocpirpc"
	"google.golang.org/grpc"
)

func (s *OcpiService) SessionCreated(ctx context.Context, in *ocpirpc.SessionCreatedRequest, opts ...grpc.CallOption) (*ocpirpc.SessionCreatedResponse, error) {
	timerStart := time.Now()
	response, err := s.getSessionClient().SessionCreated(ctx, in, opts...)
	timerStop := time.Now()

	log.Printf("SessionCreated responded in %f seconds", timerStop.Sub(timerStart).Seconds())

	return response, err
}

func (s *OcpiService) SessionUpdated(ctx context.Context, in *ocpirpc.SessionUpdatedRequest, opts ...grpc.CallOption) (*ocpirpc.SessionUpdatedResponse, error) {
	timerStart := time.Now()
	response, err := s.getSessionClient().SessionUpdated(ctx, in, opts...)
	timerStop := time.Now()

	log.Printf("SessionUpdated responded in %f seconds", timerStop.Sub(timerStart).Seconds())

	return response, err
}

func (s *OcpiService) getSessionClient() ocpirpc.SessionServiceClient {
	if s.sessionClient == nil {
		client := ocpirpc.NewSessionServiceClient(s.clientConn)
		s.sessionClient = &client
	}

	return *s.sessionClient
}
