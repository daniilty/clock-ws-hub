package core

import (
	"context"
	"fmt"

	"github.com/daniilty/tinkoff-invest-diff/records"
)

func (s *ServiceImpl) GetPortfolioDiff(ctx context.Context, id string) (string, error) {
	portfolio, err := s.tinkoffClient.Portfolio(ctx, id)
	if err != nil {
		return "", fmt.Errorf("get portfolio: %w", err)
	}

	diff := records.GetDiffString(portfolio.Positions)

	return diff, nil
}
