package xkcd

import (
	"context"
	"github.com/ferryvg/go-feeds-aggregator/internal/domain"
)

type Service struct {
	cache Cache
}

func (s *Service) Load(ctx context.Context, limit int) ([]domain.FeedItem, error) {
	return nil, nil
}
