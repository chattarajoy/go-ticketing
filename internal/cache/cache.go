package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

type Type string

const InMemoryCache Type = "inMemory"

type Cache interface {
	Set(key string, object interface{}, duration time.Duration)
	Get(key string) (interface{}, bool)
	Delete(key string)
}

func NewCache(cacheType Type) Cache {
	c := cache.New(5*time.Minute, 10*time.Minute)
	if cacheType == "inMemory" {
		// Create a cache with a default expiration time of 5 minutes, and which
		// purges expired items every 10 minutes
		return c
	}
	return c
}
