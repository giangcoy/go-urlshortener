package store

import (
	"time"

	"github.com/giangcoy/go-urlshortener/internal/store/memory"
)

type Store interface {
	Get(string) (string, error)
	Set(string, string, time.Duration) error
}

func NewMemory() Store {
	return memory.New()
}
