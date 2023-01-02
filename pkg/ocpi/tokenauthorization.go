package ocpi

import (
	"context"
	"log"
	"time"

	"github.com/satimoto/go-ocpi/ocpirpc"
	"google.golang.org/grpc"
)

func (s *OcpiService) UpdateTokenAuthorization(ctx context.Context, in *ocpirpc.UpdateTokenAuthorizationRequest, opts ...grpc.CallOption) (*ocpirpc.UpdateTokenAuthorizationResponse, error) {
	timerStart := time.Now()
	response, err := s.getTokenAuthorizationClient().UpdateTokenAuthorization(ctx, in, opts...)
	timerStop := time.Now()

	log.Printf("UpdateTokenAuthorization responded in %f seconds", timerStop.Sub(timerStart).Seconds())

	return response, err
}

func (s *OcpiService) getTokenAuthorizationClient() ocpirpc.TokenAuthorizationServiceClient {
	if s.tokenAuthorizationClient == nil {
		client := ocpirpc.NewTokenAuthorizationServiceClient(s.clientConn)
		s.tokenAuthorizationClient = &client
	}

	return *s.tokenAuthorizationClient
}
