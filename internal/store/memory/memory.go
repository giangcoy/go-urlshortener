package memory

import (
	"sync"
	"time"
)

type Memory struct {
	mtx sync.RWMutex
	s   map[string]string
}

func New() *Memory {
	return &Memory{
		mtx: sync.RWMutex{},
		s:   map[string]string{},
	}
}
func (m *Memory) Get(key string) (string, error) {
	m.mtx.RLock()
	defer m.mtx.RUnlock()
	value, ok := m.s[key]
	if ok {
		return value, nil
	}
	return "", nil
}
func (m *Memory) Set(key string, value string, ttl time.Duration) error {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	m.s[key] = value
	return nil

}
