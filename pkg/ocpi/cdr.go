package ocpi

import (
	"context"
	"log"
	"time"

	"github.com/satimoto/go-ocpi/ocpirpc"
	"google.golang.org/grpc"
)

func (s *OcpiService) CdrCreated(ctx context.Context, in *ocpirpc.CdrCreatedRequest, opts ...grpc.CallOption) (*ocpirpc.CdrCreatedResponse, error) {
	timerStart := time.Now()
	response, err := s.getCdrClient().CdrCreated(ctx, in, opts...)
	timerStop := time.Now()

	log.Printf("CdrCreated responded in %f seconds", timerStop.Sub(timerStart).Seconds())

	return response, err
}

func (s *OcpiService) getCdrClient() ocpirpc.CdrServiceClient {
	if s.cdrClient == nil {
		client := ocpirpc.NewCdrServiceClient(s.clientConn)
		s.cdrClient = &client
	}

	return *s.cdrClient
}
