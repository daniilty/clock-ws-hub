package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/daniilty/clock-ws-hub/internal/core"
	"github.com/daniilty/clock-ws-hub/internal/server"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	invest "github.com/TinkoffCreditSystems/invest-openapi-go-sdk"
	schema "github.com/daniilty/weather-gateway-schema"
)

func run() error {
	cfg, err := loadEnvConfig()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithCancel(context.Background())

	conn, err := grpc.DialContext(ctx, cfg.weatherGRPCAddr, grpc.WithInsecure())
	if err != nil {
		cancel()

		return err
	}

	weatherClient := schema.NewGismeteoWeatherGatewayClient(conn)
	tinkoffClient := invest.NewRestClient(cfg.tinkoffAPIKey)

	service := core.NewServiceImpl(weatherClient, tinkoffClient)

	loggerCfg := zap.NewProductionConfig()

	logger, err := loggerCfg.Build()
	if err != nil {
		cancel()

		return err
	}

	sugaredLogger := logger.Sugar()

	wsServer := server.NewWS(service, sugaredLogger, 1*time.Minute, cfg.httpAddr, cfg.tinkoffAccountID)

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func(ctx context.Context) {
		wsServer.Run(ctx)
		wg.Done()
	}(ctx)

	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-termChan

	cancel()
	wg.Wait()

	return nil
}

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}
