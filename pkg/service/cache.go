package service

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	cache "memcache/api/proto"
)

type CacheServer struct {
	memCache Storage
}

func NewCacheServer(st Storage) *CacheServer {
	return &CacheServer{memCache: st}
}

func (c *CacheServer) Set(_ context.Context, item *cache.Item) (*empty.Empty, error) {
	err := c.memCache.Set(item.Key, item.Value)
	if err != nil {
		return nil, err
	}
	out := new(empty.Empty)
	return out, err
}

func (c *CacheServer) Get(_ context.Context, key *cache.Key) (*cache.Item, error) {
	val, err := c.memCache.Get(key.Key)
	if err != nil {
		return nil, err
	}
	item := cache.Item{Key: key.Key, Value: val}
	return &item, err
}

func (c *CacheServer) Delete(_ context.Context, key *cache.Key) (*empty.Empty, error) {
	err := c.memCache.Delete(key.Key)
	if err != nil {
		return nil, err
	}
	out := new(empty.Empty)
	return out, err
}
