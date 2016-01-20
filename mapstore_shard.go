package mapstore

import (
	"sync"

	"github.com/spaolacci/murmur3"
)

var _ Store = (*StoreShard)(nil)

type shard struct {
	m sync.RWMutex
	s map[string]interface{}
}

type StoreShard struct {
	m           sync.RWMutex
	shards      []*shard
	shardsCount int
}

func newShard() *shard {
	return &shard{
		s: make(map[string]interface{}),
	}
}

func newStoreShard(shardsCount int) *StoreShard {
	store := &StoreShard{}

	store.shardsCount = shardsCount
	store.shards = make([]*shard, shardsCount)

	for k := range store.shards {
		store.shards[k] = newShard()
	}

	return store
}

func (s *StoreShard) getShard(key string) *shard {
	s.m.RLock()
	defer s.m.RUnlock()
	num := int(murmur3.Sum64([]byte(key))>>1) % s.shardsCount
	return s.shards[num]
}

func (s *StoreShard) Get(key string, defaultValue interface{}) (interface{}, bool) {
	shard := s.getShard(key)
	shard.m.RLock()
	value, ok := shard.s[key]
	shard.m.RUnlock()

	if !ok {
		return defaultValue, false
	}
	return value, true
}

func (s *StoreShard) GetOrSet(key string, defaultValue interface{}) (interface{}, bool) {
	shard := s.getShard(key)
	shard.m.Lock()
	value, ok := shard.s[key]
	if !ok {
		shard.s[key] = defaultValue
		value = defaultValue
	}
	shard.m.Unlock()
	return value, ok
}

func (s *StoreShard) Delete(key string) (bool) {
	shard := s.getShard(key)
	shard.m.Lock()
	_, ok := shard.s[key]
	delete(shard.s, key)
	shard.m.Unlock()
	return ok
}

func (s *StoreShard) Set(key string, value interface{}) {
	shard := s.getShard(key)
	shard.m.Lock()
	shard.s[key] = value
	shard.m.Unlock()
}

func (s *StoreShard) Load(entries chan Entry) {
	s.m.Lock()
	defer s.m.Unlock()
	for entry := range entries {
		shard := s.getShard(entry.Key)
		shard.s[entry.Key] = entry.Value
	}
}

func (s *StoreShard) Save(entries chan<- Entry) {
	s.m.RLock()
	defer s.m.RUnlock()
	for _, shard := range s.shards {
		for key, value := range shard.s {
			entries <- Entry{key, value}
		}
	}
}

func (s *StoreShard) ShardStats() []int {
	result := make([]int, s.shardsCount)
	for i := 0; i < s.shardsCount; i++ {
		shard := s.shards[i]
		shard.m.RLock()
		result[i] = len(shard.s)
		shard.m.RUnlock()
	}
	return result
}

func (s *StoreShard) Len() int {
	return 0
}
