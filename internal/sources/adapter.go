package sources

import (
	"context"
	"github.com/ferryvg/go-feeds-aggregator/internal/domain"
)

type SourceAdapter interface {
	Load(ctx context.Context, limit int) ([]domain.FeedItem, error)
}
