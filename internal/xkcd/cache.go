package xkcd

import "errors"

var NotFoundErr = errors.New("comics item does not exists in cache")

type Cache interface {
	// Get returns comics item by id. If not exists returns NotFoundErr
	Get(id int) (*SourceItem, error)

	// Set store comics item in cache
	Set(item *SourceItem) error
}
