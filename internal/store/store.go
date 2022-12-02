package store

import (
	"time"

	"github.com/giangcoy/go-urlshortener/internal/store/memory"
	"github.com/giangcoy/go-urlshortener/internal/store/redis"
)

type Store interface {
	Get(string) (string, error)
	Set(string, string, time.Duration) error
}

func NewMemory() Store {
	return memory.New()
}
func NewRedis(addr, prefix string) Store {
	return redis.New(addr, prefix)
}
