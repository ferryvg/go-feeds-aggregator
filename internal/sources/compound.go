package sources

import (
	"context"

	"github.com/ferryvg/go-feeds-aggregator/internal/domain"
	"golang.org/x/sync/errgroup"
)

type CompoundAdapter struct {
	adapters []SourceAdapter
}

func NewCompoundAdapter(adapters ...SourceAdapter) *CompoundAdapter {
	return &CompoundAdapter{
		adapters: adapters,
	}
}

func (c *CompoundAdapter) Init() error {
	for _, adapter := range c.adapters {
		if err := adapter.Init(); err != nil {
			return err
		}
	}

	return nil
}

func (c *CompoundAdapter) Stop() {
	for _, adapter := range c.adapters {
		adapter.Stop()
	}
}

func (c *CompoundAdapter) Load(ctx context.Context, limit int) ([]*domain.FeedItem, error) {
	if limit <= 0 {
		return nil, nil
	}

	perSource := limit / len(c.adapters)

	group, ctx := errgroup.WithContext(ctx)

	stack := NewStack()

	for _, adapter := range c.adapters {
		current := adapter
		group.Go(func() error {
			items, err := current.Load(ctx, perSource)
			if err != nil {
				return err
			}

			stack.Add(items...)

			return nil
		})
	}

	if err := group.Wait(); err != nil {
		return nil, err
	}

	stack.SortByDate()

	return stack.All(), nil
}
