package cache

import (
	"github.com/redis/go-redis/v9"
	"github.com/skhanal5/redis-tryout/internal"
)

type CacheManager interface {
	LobbyManager
}

type Cache struct {
	Client redis.Client
}

func NewCache(options internal.Options) Cache {
	client := redis.NewClient(&redis.Options{
		Addr: options.RedisAddr,
	})
	return Cache{
		Client: *client,
	}
}