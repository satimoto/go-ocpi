package rpc

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"sync"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	"github.com/satimoto/go-ocpi/internal/rpc/command"
	"github.com/satimoto/go-ocpi/internal/rpc/credential"
	"github.com/satimoto/go-ocpi/internal/rpc/rpc"
	"github.com/satimoto/go-ocpi/internal/rpc/token"
	"github.com/satimoto/go-ocpi/internal/service"
	ocpiSync "github.com/satimoto/go-ocpi/internal/sync"
	"github.com/satimoto/go-ocpi/internal/transportation"
	"github.com/satimoto/go-ocpi/ocpirpc"
	"google.golang.org/grpc"
)

type Rpc interface {
	StartRpc(context.Context, *sync.WaitGroup)
}

type RpcService struct {
	RepositoryService     *db.RepositoryService
	OcpiService           *transportation.OcpiService
	SyncService           *ocpiSync.SyncService
	Server                *grpc.Server
	RpcCommandResolver    *command.RpcCommandResolver
	RpcCredentialResolver *credential.RpcCredentialResolver
	RpcResolver           *rpc.RpcResolver
	RpcTokenResolver      *token.RpcTokenResolver
}

func NewRpc(repositoryService *db.RepositoryService, services *service.ServiceResolver) Rpc {
	return &RpcService{
		RepositoryService:     repositoryService,
		OcpiService:           services.OcpiService,
		SyncService:           services.SyncService,
		Server:                grpc.NewServer(),
		RpcCommandResolver:    command.NewResolver(repositoryService, services),
		RpcCredentialResolver: credential.NewResolver(repositoryService, services),
		RpcResolver:           rpc.NewResolver(repositoryService),
		RpcTokenResolver:      token.NewResolver(repositoryService, services),
	}
}

func (rs *RpcService) StartRpc(ctx context.Context, waitGroup *sync.WaitGroup) {
	log.Printf("Starting Rpc service")
	waitGroup.Add(1)

	go rs.listenAndServe()

	go func() {
		<-ctx.Done()
		log.Printf("Shutting down Rpc service")

		rs.shutdown()

		log.Printf("Rpc service shut down")
		waitGroup.Done()
	}()
}

func (rs *RpcService) listenAndServe() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("RPC_PORT")))
	util.PanicOnError("OCPI274", "Error creating network address", err)

	ocpirpc.RegisterCommandServiceServer(rs.Server, rs.RpcCommandResolver)
	ocpirpc.RegisterCredentialServiceServer(rs.Server, rs.RpcCredentialResolver)
	ocpirpc.RegisterRpcServiceServer(rs.Server, rs.RpcResolver)
	ocpirpc.RegisterTokenServiceServer(rs.Server, rs.RpcTokenResolver)

	err = rs.Server.Serve(listener)

	if err != nil {
		util.LogOnError("OCPI278", "Error in Rpc service", err)
	}
}

func (rs *RpcService) shutdown() {
	rs.Server.GracefulStop()
}
