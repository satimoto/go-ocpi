package ocpi

import (
	"context"
	"os"

	"github.com/satimoto/go-datastore/util"
	"github.com/satimoto/go-ocpi-api/ocpirpc"
	"google.golang.org/grpc"
)

type Ocpi interface {
	StopSession(ctx context.Context, in *ocpirpc.StopSessionRequest, opts ...grpc.CallOption) (*ocpirpc.StopSessionResponse, error)
	UpdateTokens(ctx context.Context, in *ocpirpc.UpdateTokensRequest, opts ...grpc.CallOption) (*ocpirpc.UpdateTokensResponse, error)
}

type OcpiService struct {
	clientConn       *grpc.ClientConn
	commandClient    *ocpirpc.CommandServiceClient
	credentialClient *ocpirpc.CredentialServiceClient
	tokenClient      *ocpirpc.TokenServiceClient
}

func NewService() Ocpi {
	clientConn, err := grpc.Dial(os.Getenv("OCPI_RPC_ADDRESS"), grpc.WithInsecure())
	util.PanicOnError("OCPI001", "Error connecting to OCPI RPC address", err)

	return &OcpiService{
		clientConn: clientConn,
	}
}
