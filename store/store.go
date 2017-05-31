// Concurrent storage.
package store

import (
	"context"
)

type item struct {
	key         string
	value       []byte
	contentType string
}

func (i *item) Value() []byte {
	return i.value
}

func (i *item) ContentType() string {
	return i.contentType
}

type store struct {
	setChannel chan Item
	hashmap    map[string]item
}

func (s *store) Set(item Item) {
	s.setChannel <- item
}

func (s *store) setWorker(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case item := <-s.setChannel:
			s.hashmap[item.key] = item
		}
	}
}

func (s *store) Get(key string) Item {
	it, ok := s.hashmap[key]
	if !ok {
		return item{}
	}
	return it
}

func New(ctx context.Context) *store {
	setCh := make(chan Item)
	s := &store{
		setChannel: setCh,
		hashmap:    make(map[string]Item),
	}
	go s.setWorker(ctx)
	return s
}

func NewItem(key, contentType string, body []byte) *item {
	return &item{
		key:         key,
		value:       body,
		contentType: contentType,
	}
}
