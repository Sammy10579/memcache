package service

import (
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes/empty"
	cache "memcache/api/proto"
	"sync"
)

type CacheServer struct {
	mx sync.RWMutex
	m  map[string]string
}

func (c *CacheServer) Get(ctx context.Context, key *cache.Key) (*cache.Item, error) {
	c.mx.RLock()
	defer c.mx.RUnlock()

	_, ok := c.m[key.Key]
	if ok != true {
		err := errors.New("no value for this key")
		return nil, err
	}
	a := cache.Item{Key: key.Key, Value: c.m[key.Key]}
	return &a, nil
}

func (c *CacheServer) Set(ctx context.Context, item *cache.Item) (*empty.Empty, error) {
	c.mx.RLock()
	defer c.mx.RLock()

	if c.m == nil {
		c.m = make(map[string]string)
	}

	c.m[item.Key] = item.Value
	out := new(empty.Empty)
	return out, nil
}

func (c *CacheServer) Delete(ctx context.Context, key *cache.Key) (*empty.Empty, error) {
	c.mx.RLock()
	defer c.mx.RUnlock()

	delete(c.m, key.Key)
	out := new(empty.Empty)
	return out, nil
}
