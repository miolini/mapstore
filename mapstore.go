package mapstore

import (
	"sync"

	"github.com/spaolacci/murmur3"
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
	shard       *_shard
	singleShard bool
}

func NewWithSize(shardsCount int) *Store {
	s := new(Store)
	s.shardsCount = shardsCount
	if shardsCount == 1 {
		s.shard = newShard()
		s.singleShard = true
	} else {
		s.shards = make([]*_shard, s.shardsCount)
		for i := 0; i < shardsCount; i++ {
			s.shards[i] = newShard()
		}
	}
	return s
}

func New() *Store {
	return NewWithSize(defaultShardsCount)
}

func (s *Store) getShard(key string) *_shard {
	s.mutex.RLock()

	var shard *_shard

	if s.singleShard {
		shard = s.shard
	} else {
		num := int(murmur3.Sum64([]byte(key))>>1) % s.shardsCount
		shard = s.shards[num]
	}

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
	if s.singleShard {
		s.shard.mutex.RLock()
		result := []int{len(s.shard.data)}
		s.shard.mutex.RUnlock()
		return result
	}

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

	for entry := range entries {
		shard := s.getShard(entry.Key)
		shard.data[entry.Key] = entry.Value
	}

	s.mutex.Unlock()
}

func (s *Store) Save(entries chan<- Entry) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	if s.singleShard {
		for key, value := range s.shard.data {
			entries <- Entry{key, value}
		}
	} else {
		for _, shard := range s.shards {
			for key, value := range shard.data {
				entries <- Entry{key, value}
			}
		}
	}
}
