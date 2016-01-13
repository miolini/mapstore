package mapstore

import (
	"sync"
	"github.com/spaolacci/murmur3"
)

const shards = 1024

type _shard struct {
	data map[string]interface{}
	mutex sync.RWMutex
}

func newShard() *_shard {
	s := new(_shard)
	s.data = make(map[string]interface{})
	return s
}

type Store struct {
	shards []*_shard
}

func New() *Store {
	s := new(Store)
	s.shards = make([]*_shard, shards)
	for i:=0;i<shards;i++ {
		s.shards[i] = newShard()
	}
	return s
}

func (s *Store) getShard(key string) *_shard {
	num := murmur3.Sum64([]byte(key)) % shards
	return s.shards[num]
}

func (s *Store) Set(key string, value interface{}) {
	shard := s.getShard(key)
	shard.mutex.Lock()
	shard.data[key] = value
	shard.mutex.Unlock()
}

func (s *Store) Get(key string, defaultValue interface{}) (interface{}, bool) {
	shard := s.getShard(key)
	shard.mutex.RLock()
	result, ok := shard.data[key]
	shard.mutex.RUnlock()
	if !ok {
		result = defaultValue
	}
	return result, ok
}