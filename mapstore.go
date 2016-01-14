package mapstore

import (
	"github.com/spaolacci/murmur3"
	"sync"
)

const defaultShardsCount = 1024

type _shard struct {
	data  map[string]interface{}
	mutex sync.RWMutex
}

func newShard() *_shard {
	s := new(_shard)
	s.data = make(map[string]interface{})
	return s
}

type Store struct {
	shards      []*_shard
	shardsCount int
	mutex       sync.RWMutex
}

func NewWithSize(shardsCount int) *Store {
	s := new(Store)
	s.mutex.Lock()
	s.shardsCount = shardsCount
	s.shards = make([]*_shard, s.shardsCount)
	for i := 0; i < shardsCount; i++ {
		s.shards[i] = newShard()
	}
	s.mutex.Unlock()
	return s
}

func New() *Store {
	return NewWithSize(defaultShardsCount)
}

func (s *Store) getShard(key string) *_shard {
	s.mutex.RLock()
	num := int(murmur3.Sum64([]byte(key))>>1) % s.shardsCount
	shard := s.shards[num]
	s.mutex.RUnlock()
	return shard
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

func (s *Store) ShardStats() []int {
	result := make([]int, s.shardsCount)
	for i := 0; i < s.shardsCount; i++ {
		shard := s.shards[i]
		shard.mutex.RLock()
		result[i] = len(shard.data)
		shard.mutex.RUnlock()
	}
	return result
}

type Entry struct {
	Key   string
	Value interface{}
}

func (s *Store) Load(entries chan Entry) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	for entry := range entries {
		shard := s.getShard(entry.Key)
		shard.data[entry.Key] = entry.Value
	}
}

func (s *Store) Save(entries chan<- Entry) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	for _, shard := range s.shards {
		for key, value := range shard.data {
			entries <- Entry{key, value}
		}
	}
}
