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

func (s *StoreSingle) GetOrSet(key string, defaultValue interface{}) (interface{}, bool) {
	s.m.Lock()
	value, ok := s.s[key]
	if !ok {
		s.s[key] = defaultValue
		value = defaultValue
	}
	s.m.Unlock()
	return value, ok
}

func (s *StoreSingle) GetOrSetFunc(key string, newFn func() interface{}) (interface{}, bool) {
	s.m.Lock()
	value, ok := s.s[key]
	if !ok {
		value = newFn()
		s.s[key] = value
	}
	s.m.Unlock()
	return value, ok
}

func (s *StoreSingle) Delete(key string) bool {
	s.m.Lock()
	_, ok := s.s[key]
	delete(s.s, key)
	s.m.Unlock()
	return ok
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
	return []int{s.Len()}
}

func (s *StoreSingle) Len() int {
	s.m.RLock()
	storeLen := len(s.s)
	s.m.RUnlock()
	return storeLen
}
