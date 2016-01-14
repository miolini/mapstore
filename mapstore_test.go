package mapstore

import (
	"testing"
	"sync"
	"fmt"
	"math/rand"
)

func BenchmarkRead50Write50(b *testing.B) {
	keys := make([]string, 10000)
	for i:=0;i<len(keys);i++ {
		keys[i] = fmt.Sprintf("%d", i)
	}
	store := New()
	wg := sync.WaitGroup{}
	wg.Add(2)
	b.ResetTimer()
	go testWrites(store, keys, b.N, &wg)
	go testReads(store, keys, b.N, &wg)
	wg.Wait()
	b.Logf("stat: %v", store.ShardStats())
}

func testWrites(s *Store, keys []string, num int, wg *sync.WaitGroup) {
	defer wg.Done()
	lenKeys := len(keys)
	for i:=0;i<num;i++ {
		key := keys[rand.Int() % lenKeys]
		s.Set(key, key)
	}
}

func testReads(s *Store, keys []string, num int, wg *sync.WaitGroup) {
	defer wg.Done()
	lenKeys := len(keys)
	for i:=0;i<num;i++ {
		key := keys[rand.Int() % lenKeys]
		s.Get(key, key)
	}
}