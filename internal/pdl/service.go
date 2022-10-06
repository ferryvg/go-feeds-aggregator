package pdl

import (
	"context"
	"github.com/ferryvg/go-feeds-aggregator/internal/domain"
)

type Service struct {
}

func (s *Service) Load(ctx context.Context, limit int) ([]domain.FeedItem, error) {
	return nil, nil
}
