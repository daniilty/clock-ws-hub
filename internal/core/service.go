package core

import (
	"context"

	invest "github.com/TinkoffCreditSystems/invest-openapi-go-sdk"
	"github.com/daniilty/clock-ws-hub/internal/pb"
)

var _ Service = (*ServiceImpl)(nil)

type Service interface {
	GetPortfolioDiff(context.Context, string) (string, error)
	GetWeather(context.Context) (string, error)
}

type ServiceImpl struct {
	weatherClient pb.GismeteoWeatherGatewayClient
	tinkoffClient *invest.RestClient
}

func NewServiceImpl(weatherClient pb.GismeteoWeatherGatewayClient, tinkoffClient *invest.RestClient) *ServiceImpl {
	return &ServiceImpl{
		weatherClient: weatherClient,
		tinkoffClient: tinkoffClient,
	}
}
