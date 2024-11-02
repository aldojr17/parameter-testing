package cache

import (
	"encoding/json"
	"parameter-testing/domain"
	log "parameter-testing/logger"
	"time"

	"github.com/go-redis/redis"
)

type Cache struct {
	redis *redis.Client
	ttl   time.Duration
}

func NewCache(redis *redis.Client, ttl time.Duration) *Cache {
	return &Cache{
		redis: redis,
		ttl:   ttl,
	}
}

func (c *Cache) getKey() string {
	return "<cache key>"
}

func (c *Cache) Set(data *domain.StructName) error {
	val, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if err := c.redis.Set(c.getKey(), val, c.ttl).Err(); err != nil {
		return err
	}

	log.Info("cache created", "key", c.getKey(), "ttl", c.ttl)

	return nil
}

func (c *Cache) Get() (*domain.StructName, error) {
	response := new(domain.StructName)

	val, err := c.redis.Get(c.getKey()).Bytes()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(val, &response)
	if err != nil {
		return nil, err
	}

	log.Info("load data from cache")

	return response, nil
}
