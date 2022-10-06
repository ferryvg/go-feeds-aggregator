package xkcd

import (
	"errors"
	"sync"
)

var NotFoundErr = errors.New("comics item does not exists in cache")

type Cache interface {
	// Get returns comics item by id. If not exists returns NotFoundErr
	Get(id int) (*SourceItem, error)

	// Set store comics item in cache
	Set(item *SourceItem)
}

type cacheImpl struct {
	data map[int]*SourceItem
	mu   *sync.RWMutex
}

func NewCache() Cache {
	return &cacheImpl{
		data: make(map[int]*SourceItem),
		mu:   new(sync.RWMutex),
	}
}

func (c *cacheImpl) Get(id int) (*SourceItem, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, exists := c.data[id]
	if !exists {
		return nil, NotFoundErr
	}

	return item, nil
}

func (c *cacheImpl) Set(item *SourceItem) {
	c.mu.Lock()
	c.data[item.Id] = item
	c.mu.Unlock()
}
