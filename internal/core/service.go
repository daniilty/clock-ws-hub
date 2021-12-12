package core

import (
	"context"

	invest "github.com/TinkoffCreditSystems/invest-openapi-go-sdk"
	schema "github.com/daniilty/weather-gateway-schema"
)

var _ Service = (*ServiceImpl)(nil)

type Service interface {
	GetPortfolioDiff(context.Context, string) (string, error)
	GetWeather(context.Context) (string, error)
}

type ServiceImpl struct {
	weatherClient schema.GismeteoWeatherGatewayClient
	tinkoffClient *invest.RestClient
}

func NewServiceImpl(weatherClient schema.GismeteoWeatherGatewayClient, tinkoffClient *invest.RestClient) *ServiceImpl {
	return &ServiceImpl{
		weatherClient: weatherClient,
		tinkoffClient: tinkoffClient,
	}
}
