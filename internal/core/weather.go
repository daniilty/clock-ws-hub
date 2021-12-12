package core

import (
	"context"
	"fmt"

	schema "github.com/daniilty/weather-gateway-schema"
)

func (s *ServiceImpl) GetWeather(ctx context.Context) (string, error) {
	resp, err := s.weatherClient.GetWeather(ctx, &schema.Empty{})
	if err != nil {
		return "", fmt.Errorf("get weather: %w", err)
	}

	return resp.GetInfo(), nil
}
