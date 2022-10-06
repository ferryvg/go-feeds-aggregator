package sources

import (
	"context"
	"github.com/ferryvg/go-feeds-aggregator/internal/domain"
	"golang.org/x/sync/errgroup"
	"sync"
)

type Stack struct {
	items []domain.FeedItem
	mu    *sync.RWMutex
}

func (s *Stack) Add(items ...domain.FeedItem) {
	s.mu.Lock()

	s.items = append(s.items, items...)

	s.mu.Unlock()
}

func (s *Stack) All() []domain.FeedItem {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.items
}

type CompoundAdapter struct {
	adapters []SourceAdapter
}

func (c *CompoundAdapter) Load(ctx context.Context, limit int) ([]domain.FeedItem, error) {
	perSource := limit / len(c.adapters)

	group, ctx := errgroup.WithContext(ctx)

	stack := new(Stack)

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

	// TODO: sort stack

	return stack.All(), group.Wait()
}
