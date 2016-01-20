package mapstore

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
)

func BenchmarkShard1Read4Write4(b *testing.B) {
	bench(1, 4, 4, 1000000, b)
}
func BenchmarkShard10Read4Write4(b *testing.B) {
	bench(10, 4, 4, 1000000, b)
}
func BenchmarkShard100Read4Write4(b *testing.B) {
	bench(100, 4, 4, 1000000, b)
}
func BenchmarkShard1000Read4Write4(b *testing.B) {
	bench(1000, 4, 4, 1000000, b)
}
func BenchmarkShard1Read8Write2(b *testing.B) {
	bench(1, 8, 2, 1000000, b)
}
func BenchmarkShard10Read8Write2(b *testing.B) {
	bench(10, 8, 2, 1000000, b)
}
func BenchmarkShard100Read8Write2(b *testing.B) {
	bench(100, 8, 2, 1000000, b)
}
func BenchmarkShard1000Read8Write2(b *testing.B) {
	bench(1000, 8, 2, 1000000, b)
}
func BenchmarkShard1Read2Write8(b *testing.B) {
	bench(1, 2, 8, 1000000, b)
}
func BenchmarkShard10Read2Write8(b *testing.B) {
	bench(10, 2, 8, 1000000, b)
}
func BenchmarkShard100Read2Write8(b *testing.B) {
	bench(100, 2, 8, 1000000, b)
}
func BenchmarkShard1000Read2Write8(b *testing.B) {
	bench(1000, 2, 8, 1000000, b)
}
func BenchmarkShard1Read8Write0(b *testing.B) {
	bench(1, 8, 0, 1000000, b)
}
func BenchmarkShard10Read8Write0(b *testing.B) {
	bench(10, 8, 0, 1000000, b)
}
func BenchmarkShard100Read8Write0(b *testing.B) {
	bench(100, 8, 0, 1000000, b)
}
func BenchmarkShard1000Read8Write0(b *testing.B) {
	bench(1000, 8, 0, 1000000, b)
}

func TestShardsStat(t *testing.T) {
	keys := genKeys(1e6)
	shards := 4
	store := NewWithSize(shards)
	for i := 0; i < 1e7; i++ {
		key := keys[i%len(keys)]
		store.Set(key, key)
	}
	t.Logf("stats: %v", store.ShardStats())
}

func genKeys(count int) []string {
	keys := make([]string, count)
	for i := 0; i < len(keys); i++ {
		keys[i] = fmt.Sprintf("%d", i)
	}
	return keys
}

func bench(shards, readThreads, writeThreads int, keysCount int, b *testing.B) {
	keys := genKeys(keysCount)
	store := NewWithSize(shards)
	wg := sync.WaitGroup{}
	b.ResetTimer()
	wg.Add(readThreads + writeThreads)
	for i := 0; i < writeThreads; i++ {
		go testWrites(store, keys, b.N, &wg)
	}
	for i := 0; i < readThreads; i++ {
		go testReads(store, keys, b.N, &wg)
	}
	wg.Wait()
}

func testWrites(s Store, keys []string, num int, wg *sync.WaitGroup) {
	defer wg.Done()
	lenKeys := len(keys)
	for i := 0; i < num; i++ {
		key := keys[rand.Int()%lenKeys]
		s.Set(key, key)
	}
}

func testReads(s Store, keys []string, num int, wg *sync.WaitGroup) {
	defer wg.Done()
	lenKeys := len(keys)
	for i := 0; i < num; i++ {
		key := keys[rand.Int()%lenKeys]
		s.Get(key, key)
	}
}
