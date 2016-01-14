package mapstore

import (
	"sync"
)

var _ Store = (*StoreSingle)(nil)

type StoreSingle struct {
	m sync.RWMutex
	s map[string]interface{}
}

func newStoreSingle() *StoreSingle {
	return &StoreSingle{
		s: make(map[string]interface{}),
	}
}

func (s *StoreSingle) Get(key string, defaultValue interface{}) (interface{}, bool) {
	s.m.RLock()
	value, ok := s.s[key]
	s.m.RUnlock()
	if !ok {
		return defaultValue, false
	}
	return value, true
}

func (s *StoreSingle) Set(key string, value interface{}) {
	s.m.Lock()
	s.s[key] = value
	s.m.Unlock()
}

func (s *StoreSingle) Load(entries chan Entry) {
	for entry := range entries {
		s.m.Lock()
		s.s[entry.Key] = entry.Value
		s.m.Unlock()
	}
}

func (s *StoreSingle) Save(entries chan<- Entry) {
	s.m.RLock()
	defer s.m.Unlock()
	for k, v := range s.s {
		entries <- Entry{k, v}
	}
}

func (s *StoreSingle) ShardStats() []int {
	return nil
}

func (s *StoreSingle) Len() int {
	s.m.RLock()
	defer s.m.RUnlock()
	return len(s.s)
}
