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
	"github.com/satimoto/go-ocpi-api/internal/rpc/command"
	"github.com/satimoto/go-ocpi-api/internal/rpc/credential"
	"github.com/satimoto/go-ocpi-api/internal/rpc/rpc"
	"github.com/satimoto/go-ocpi-api/internal/rpc/token"
	"github.com/satimoto/go-ocpi-api/ocpirpc"
	"google.golang.org/grpc"
)

type Rpc interface {
	StartRpc(context.Context, *sync.WaitGroup)
}

type RpcService struct {
	RepositoryService     *db.RepositoryService
	Server                *grpc.Server
	RpcCommandResolver    *command.RpcCommandResolver
	RpcCredentialResolver *credential.RpcCredentialResolver
	RpcResolver           *rpc.RpcResolver
	RpcTokenResolver      *token.RpcTokenResolver
}

func NewRpc(d *sql.DB) Rpc {
	repositoryService := db.NewRepositoryService(d)

	return &RpcService{
		RepositoryService:     repositoryService,
		Server:                grpc.NewServer(),
		RpcCommandResolver:    command.NewResolver(repositoryService),
		RpcCredentialResolver: credential.NewResolver(repositoryService),
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

	if err != nil {
		log.Printf("Error creating network address: %v", err)
	}

	ocpirpc.RegisterCommandServiceServer(rs.Server, rs.RpcCommandResolver)
	ocpirpc.RegisterCredentialServiceServer(rs.Server, rs.RpcCredentialResolver)
	ocpirpc.RegisterRpcServiceServer(rs.Server, rs.RpcResolver)
	ocpirpc.RegisterTokenServiceServer(rs.Server, rs.RpcTokenResolver)

	err = rs.Server.Serve(listener)

	if err != nil {
		log.Printf("Error in Rpc service: %v", err)
	}
}

func (rs *RpcService) shutdown() {
	rs.Server.GracefulStop()
}
