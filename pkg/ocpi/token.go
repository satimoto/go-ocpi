package ocpi

import (
	"context"
	"log"
	"time"

	"github.com/satimoto/go-ocpi/ocpirpc"
	"google.golang.org/grpc"
)

func (s *OcpiService) CreateToken(ctx context.Context, in *ocpirpc.CreateTokenRequest, opts ...grpc.CallOption) (*ocpirpc.CreateTokenResponse, error) {
	timerStart := time.Now()
	response, err := s.getTokenClient().CreateToken(ctx, in, opts...)
	timerStop := time.Now()

	log.Printf("CreateToken responded in %f seconds", timerStop.Sub(timerStart).Seconds())

	return response, err
}

func (s *OcpiService) UpdateTokens(ctx context.Context, in *ocpirpc.UpdateTokensRequest, opts ...grpc.CallOption) (*ocpirpc.UpdateTokensResponse, error) {
	timerStart := time.Now()
	response, err := s.getTokenClient().UpdateTokens(ctx, in, opts...)
	timerStop := time.Now()

	log.Printf("UpdateTokens responded in %f seconds", timerStop.Sub(timerStart).Seconds())

	return response, err
}

func (s *OcpiService) getTokenClient() ocpirpc.TokenServiceClient {
	if s.tokenClient == nil {
		client := ocpirpc.NewTokenServiceClient(s.clientConn)
		s.tokenClient = &client
	}

	return *s.tokenClient
}
