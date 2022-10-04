package ocpi

import (
	"context"

	"github.com/satimoto/go-datastore/pkg/util"
	"github.com/satimoto/go-ocpi/ocpirpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Ocpi interface {
	TestConnection(ctx context.Context, in *ocpirpc.TestConnectionRequest, opts ...grpc.CallOption) (*ocpirpc.TestConnectionResponse, error)
	TestMessage(ctx context.Context, in *ocpirpc.TestMessageRequest, opts ...grpc.CallOption) (*ocpirpc.TestMessageResponse, error)

	CdrCreated(ctx context.Context, in *ocpirpc.CdrCreatedRequest, opts ...grpc.CallOption) (*ocpirpc.CdrCreatedResponse, error)

	CreateCredential(ctx context.Context, in *ocpirpc.CreateCredentialRequest, opts ...grpc.CallOption) (*ocpirpc.CreateCredentialResponse, error)
	RegisterCredential(ctx context.Context, in *ocpirpc.RegisterCredentialRequest, opts ...grpc.CallOption) (*ocpirpc.RegisterCredentialResponse, error)
	SyncCredential(ctx context.Context, in *ocpirpc.SyncCredentialRequest, opts ...grpc.CallOption) (*ocpirpc.SyncCredentialResponse, error)
	UnregisterCredential(ctx context.Context, in *ocpirpc.UnregisterCredentialRequest, opts ...grpc.CallOption) (*ocpirpc.UnregisterCredentialResponse, error)

	StartSession(ctx context.Context, in *ocpirpc.StartSessionRequest, opts ...grpc.CallOption) (*ocpirpc.StartSessionResponse, error)
	StopSession(ctx context.Context, in *ocpirpc.StopSessionRequest, opts ...grpc.CallOption) (*ocpirpc.StopSessionResponse, error)

	SessionCreated(ctx context.Context, in *ocpirpc.SessionCreatedRequest, opts ...grpc.CallOption) (*ocpirpc.SessionCreatedResponse, error)
	SessionUpdated(ctx context.Context, in *ocpirpc.SessionUpdatedRequest, opts ...grpc.CallOption) (*ocpirpc.SessionUpdatedResponse, error)

	CreateToken(ctx context.Context, in *ocpirpc.CreateTokenRequest, opts ...grpc.CallOption) (*ocpirpc.CreateTokenResponse, error)
	UpdateTokens(ctx context.Context, in *ocpirpc.UpdateTokensRequest, opts ...grpc.CallOption) (*ocpirpc.UpdateTokensResponse, error)
}

type OcpiService struct {
	clientConn       *grpc.ClientConn
	commandClient    *ocpirpc.CommandServiceClient
	cdrClient        *ocpirpc.CdrServiceClient
	credentialClient *ocpirpc.CredentialServiceClient
	rpcClient        *ocpirpc.RpcServiceClient
	sessionClient    *ocpirpc.SessionServiceClient
	tokenClient      *ocpirpc.TokenServiceClient
}

func NewService(address string) Ocpi {
	clientConn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	util.PanicOnError("OCPI001", "Error connecting to OCPI RPC address", err)

	return &OcpiService{
		clientConn: clientConn,
	}
}
