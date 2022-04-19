package rpc

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"
	"sync"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/rpc/command"
	"github.com/satimoto/go-ocpi-api/internal/rpc/credential"
	"github.com/satimoto/go-ocpi-api/internal/rpc/token"
	"github.com/satimoto/go-ocpi-api/ocpirpc/commandrpc"
	"github.com/satimoto/go-ocpi-api/ocpirpc/credentialrpc"
	"github.com/satimoto/go-ocpi-api/ocpirpc/tokenrpc"
	"google.golang.org/grpc"
)

type Rpc interface {
	StartRpc(context.Context, *sync.WaitGroup)
}

type RpcService struct {
	*db.RepositoryService
	*grpc.Server
	*command.RpcCommandResolver
	*credential.RpcCredentialResolver
	*token.RpcTokenResolver
}

func NewRpc(d *sql.DB) Rpc {
	repositoryService := db.NewRepositoryService(d)

	return &RpcService{
		RepositoryService:     repositoryService,
		Server:                grpc.NewServer(),
		RpcCommandResolver:    command.NewResolver(repositoryService),
		RpcCredentialResolver: credential.NewResolver(repositoryService),
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

	commandrpc.RegisterCommandServiceServer(rs.Server, rs.RpcCommandResolver)
	credentialrpc.RegisterCredentialServiceServer(rs.Server, rs.RpcCredentialResolver)
	tokenrpc.RegisterTokenServiceServer(rs.Server, rs.RpcTokenResolver)

	err = rs.Server.Serve(listener)

	if err != nil {
		log.Printf("Error in Rpc service: %v", err)
	}
}

func (rs *RpcService) shutdown() {
	rs.Server.GracefulStop()
}
