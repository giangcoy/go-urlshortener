package redis

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

type Redis struct {
	rdb    *redis.Client
	prefix string
}

func New(addr, prefix string) *Redis {
	return &Redis{
		rdb:    redis.NewClient(&redis.Options{Addr: addr}),
		prefix: prefix,
	}
}
func (r *Redis) wrapKey(key string) string {
	return r.prefix + key
}
func (r *Redis) Get(key string) (string, error) {
	if r.rdb == nil {
		fmt.Println("Redis nil")
		return "", nil
	}
	val, err := r.rdb.Get(r.wrapKey(key)).Result()
	if err != nil {
		return "", nil
	}
	return val, nil
}
func (r *Redis) Set(key, value string, ttl time.Duration) error {
	if r.rdb == nil {
		fmt.Println("Redis nil")
		return nil
	}
	return r.rdb.Set(r.wrapKey(key), value, ttl).Err()

}
