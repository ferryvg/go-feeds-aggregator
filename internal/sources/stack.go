package sources

import (
	"sort"
	"sync"

	"github.com/ferryvg/go-feeds-aggregator/internal/domain"
)

type Stack struct {
	items []*domain.FeedItem
	mu    *sync.RWMutex
}

func NewStack() *Stack {
	return &Stack{
		mu: new(sync.RWMutex),
	}
}

func (s *Stack) Add(items ...*domain.FeedItem) {
	s.mu.Lock()

	s.items = append(s.items, items...)

	s.mu.Unlock()
}

func (s *Stack) All() []*domain.FeedItem {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.items
}

func (s *Stack) SortByDate() {
	if len(s.items) == 0 {
		return
	}

	sort.Slice(s.items, func(i, j int) bool {
		return s.items[i].PublishDate.Before(s.items[j].PublishDate)
	})
}
