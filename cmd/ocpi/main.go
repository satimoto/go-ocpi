package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	_ "github.com/joho/godotenv/autoload"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	"github.com/satimoto/go-ocpi/internal/metric"
	"github.com/satimoto/go-ocpi/internal/rest"
	"github.com/satimoto/go-ocpi/internal/rpc"
	"github.com/satimoto/go-ocpi/internal/service"
)

var (
	database *sql.DB

	dbHost  = os.Getenv("DB_HOST")
	dbName  = os.Getenv("DB_NAME")
	dbPass  = os.Getenv("DB_PASS")
	dbUser  = os.Getenv("DB_USER")
	sslMode = util.GetEnv("SSL_MODE", "disable")
)

func init() {
	if len(dbHost) == 0 || len(dbName) == 0 || len(dbPass) == 0 || len(dbUser) == 0 {
		log.Fatalf("Database env variables not defined")
	}

	dataSourceName := fmt.Sprintf("postgres://%s:%s@%s/%s?binary_parameters=yes&sslmode=%s", dbUser, dbPass, dbHost, dbName, sslMode)
	d, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	database = d
}

func main() {
	defer database.Close()

	log.Printf("Starting up OCPI server")
	shutdownCtx, cancelFunc := context.WithCancel(context.Background())
	waitGroup := &sync.WaitGroup{}

	repositoryService := db.NewRepositoryService(database)
	services := service.NewService(repositoryService)
	
	metricService := metric.NewMetric()
	metricService.StartMetric(shutdownCtx, waitGroup)

	restService := rest.NewRest(repositoryService, services)
	restService.StartRest(shutdownCtx, waitGroup)

	rpcService := rpc.NewRpc(repositoryService, services)
	rpcService.StartRpc(shutdownCtx, waitGroup)

	services.SyncService.StartService(shutdownCtx, waitGroup)

	sigtermChan := make(chan os.Signal)
	signal.Notify(sigtermChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigtermChan

	log.Printf("Shutting down OCPI server")

	cancelFunc()
	waitGroup.Wait()

	log.Printf("OCPI server shut down")
}
