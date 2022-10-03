package rpc

import (
	"context"
	"database/sql"
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
	opciSync "github.com/satimoto/go-ocpi/internal/sync"
	"github.com/satimoto/go-ocpi/internal/transportation"
	"github.com/satimoto/go-ocpi/ocpirpc"
	"google.golang.org/grpc"
)

type Rpc interface {
	StartRpc(context.Context, *sync.WaitGroup)
}

type RpcService struct {
	RepositoryService     *db.RepositoryService
	OcpiRequester         *transportation.OcpiRequester
	SyncService           *opciSync.SyncService
	Server                *grpc.Server
	RpcCommandResolver    *command.RpcCommandResolver
	RpcCredentialResolver *credential.RpcCredentialResolver
	RpcResolver           *rpc.RpcResolver
	RpcTokenResolver      *token.RpcTokenResolver
}

func NewRpc(d *sql.DB) Rpc {
	repositoryService := db.NewRepositoryService(d)
	ocpiRequester := transportation.NewOcpiRequester()
	syncService := opciSync.NewService(repositoryService, ocpiRequester)

	return &RpcService{
		RepositoryService:     repositoryService,
		OcpiRequester:         ocpiRequester,
		SyncService:           syncService,
		Server:                grpc.NewServer(),
		RpcCommandResolver:    command.NewResolver(repositoryService, syncService, ocpiRequester),
		RpcCredentialResolver: credential.NewResolver(repositoryService, syncService, ocpiRequester),
		RpcResolver:           rpc.NewResolver(repositoryService),
		RpcTokenResolver:      token.NewResolver(repositoryService),
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
