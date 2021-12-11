package core

import (
	"context"
	"fmt"

	"github.com/daniilty/clock-ws-hub/internal/pb"
)

func (s *ServiceImpl) GetWeather(ctx context.Context) (string, error) {
	resp, err := s.weatherClient.GetWeather(ctx, &pb.Empty{})
	if err != nil {
		return "", fmt.Errorf("get weather: %w", err)
	}

	return resp.GetInfo(), nil
}
