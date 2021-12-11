package main

import (
	"fmt"
	"os"
)

type envConfig struct {
	weatherGRPCAddr  string
	tinkoffAPIKey    string
	tinkoffAccountID string
	httpAddr         string
}

func loadEnvConfig() (*envConfig, error) {
	const (
		provideEnvErrorMsg = `please provide "%s" environment variable`

		weatherGRPCAddrEnv  = "WEATHER_GRPC_ADDR"
		tinkoffAPIKeyEnv    = "TINKOFF_API_KEY"
		tinkoffAccountIDEnv = "TINKOFF_ACCOUNT_ID"
		httpAddrEnv         = "HTTP_SERVER_ADDR"
	)

	var ok bool

	cfg := &envConfig{}

	cfg.weatherGRPCAddr, ok = os.LookupEnv(weatherGRPCAddrEnv)
	if !ok {
		return nil, fmt.Errorf(provideEnvErrorMsg, weatherGRPCAddrEnv)
	}

	cfg.tinkoffAPIKey, ok = os.LookupEnv(tinkoffAPIKeyEnv)
	if !ok {
		return nil, fmt.Errorf(provideEnvErrorMsg, tinkoffAPIKeyEnv)
	}

	cfg.tinkoffAccountID, ok = os.LookupEnv(tinkoffAccountIDEnv)
	if !ok {
		return nil, fmt.Errorf(provideEnvErrorMsg, tinkoffAccountIDEnv)
	}

	cfg.httpAddr, ok = os.LookupEnv(httpAddrEnv)
	if !ok {
		return nil, fmt.Errorf(provideEnvErrorMsg, httpAddrEnv)
	}

	return cfg, nil
}
