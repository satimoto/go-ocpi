package ocpi

import (
	"context"

	"github.com/satimoto/go-datastore/util"
	"github.com/satimoto/go-ocpi-api/ocpirpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Ocpi interface {
	CreateCredential(ctx context.Context, in *ocpirpc.CreateCredentialRequest, opts ...grpc.CallOption) (*ocpirpc.CreateCredentialResponse, error)
	RegisterCredential(ctx context.Context, in *ocpirpc.RegisterCredentialRequest, opts ...grpc.CallOption) (*ocpirpc.RegisterCredentialResponse, error)
	UnregisterCredential(ctx context.Context, in *ocpirpc.UnregisterCredentialRequest, opts ...grpc.CallOption) (*ocpirpc.UnregisterCredentialResponse, error)
	
	StartSession(ctx context.Context, in *ocpirpc.StartSessionRequest, opts ...grpc.CallOption) (*ocpirpc.StartSessionResponse, error)
	StopSession(ctx context.Context, in *ocpirpc.StopSessionRequest, opts ...grpc.CallOption) (*ocpirpc.StopSessionResponse, error)

	CreateToken(ctx context.Context, in *ocpirpc.CreateTokenRequest, opts ...grpc.CallOption) (*ocpirpc.CreateTokenResponse, error)
	UpdateTokens(ctx context.Context, in *ocpirpc.UpdateTokensRequest, opts ...grpc.CallOption) (*ocpirpc.UpdateTokensResponse, error)
}

type OcpiService struct {
	clientConn       *grpc.ClientConn
	commandClient    *ocpirpc.CommandServiceClient
	credentialClient *ocpirpc.CredentialServiceClient
	tokenClient      *ocpirpc.TokenServiceClient
}

func NewService(address string) Ocpi {
	clientConn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	util.PanicOnError("OCPI001", "Error connecting to OCPI RPC address", err)

	return &OcpiService{
		clientConn: clientConn,
	}
}
