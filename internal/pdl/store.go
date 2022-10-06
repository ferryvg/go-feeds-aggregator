package pdl

import "sync"

type Store struct {
	data []*RssItem
	mu   *sync.RWMutex
}

func NewStore() *Store {
	return &Store{
		data: nil,
		mu:   new(sync.RWMutex),
	}
}

func (s *Store) GetMulti(limit int) []*RssItem {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if limit > len(s.data) {
		return append(s.data[:0:0], s.data...)
	}

	return append(s.data[:0:limit], s.data[:limit]...)
}

func (s *Store) Update(items []*RssItem) {
	s.mu.Lock()
	s.mu.Unlock()

	s.data = append(items[:0:0], items...)
}
