package ocpi

import (
	"context"
	"log"
	"time"

	"github.com/satimoto/go-ocpi/ocpirpc"
	"google.golang.org/grpc"
)

func (s *OcpiService) CreateCredential(ctx context.Context, in *ocpirpc.CreateCredentialRequest, opts ...grpc.CallOption) (*ocpirpc.CreateCredentialResponse, error) {
	timerStart := time.Now()
	response, err := s.getCredentialClient().CreateCredential(ctx, in, opts...)
	timerStop := time.Now()

	log.Printf("CreateCredential responded in %f seconds", timerStop.Sub(timerStart).Seconds())

	return response, err
}

func (s *OcpiService) RegisterCredential(ctx context.Context, in *ocpirpc.RegisterCredentialRequest, opts ...grpc.CallOption) (*ocpirpc.RegisterCredentialResponse, error) {
	timerStart := time.Now()
	response, err := s.getCredentialClient().RegisterCredential(ctx, in, opts...)
	timerStop := time.Now()

	log.Printf("RegisterCredential responded in %f seconds", timerStop.Sub(timerStart).Seconds())

	return response, err
}

func (s *OcpiService) SyncCredential(ctx context.Context, in *ocpirpc.SyncCredentialRequest, opts ...grpc.CallOption) (*ocpirpc.SyncCredentialResponse, error) {
	timerStart := time.Now()
	response, err := s.getCredentialClient().SyncCredential(ctx, in, opts...)
	timerStop := time.Now()

	log.Printf("SyncCredential responded in %f seconds", timerStop.Sub(timerStart).Seconds())

	return response, err
}

func (s *OcpiService) UnregisterCredential(ctx context.Context, in *ocpirpc.UnregisterCredentialRequest, opts ...grpc.CallOption) (*ocpirpc.UnregisterCredentialResponse, error) {
	timerStart := time.Now()
	response, err := s.getCredentialClient().UnregisterCredential(ctx, in, opts...)
	timerStop := time.Now()

	log.Printf("UnregisterCredential responded in %f seconds", timerStop.Sub(timerStart).Seconds())

	return response, err
}

func (s *OcpiService) getCredentialClient() ocpirpc.CredentialServiceClient {
	if s.credentialClient == nil {
		client := ocpirpc.NewCredentialServiceClient(s.clientConn)
		s.credentialClient = &client
	}

	return *s.credentialClient
}
