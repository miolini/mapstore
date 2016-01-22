# mapstore
Multicore optimized in-memory data store in Go

[![Build Status](https://travis-ci.org/miolini/mapstore.svg)](https://travis-ci.org/miolini/mapstore.svg) [![Go Report Card](http://goreportcard.com/badge/miolini/mapstore)](http://goreportcard.com/report/miolini/mapstore)

## Benchmark

```bash
$ go test -v -benchmem -bench .
=== RUN   TestShardsStat
--- PASS: TestShardsStat (5.69s)
	mapstore_test.go:67: stats: [249597 249954 249518 250931]
PASS
BenchmarkShard1Read4Write4-8   	  300000	      5507 ns/op	     403 B/op	       8 allocs/op
BenchmarkShard10Read4Write4-8  	  300000	      4131 ns/op	     591 B/op	      16 allocs/op
BenchmarkShard100Read4Write4-8 	 1000000	      2316 ns/op	     384 B/op	      16 allocs/op
BenchmarkShard1000Read4Write4-8	 1000000	      2199 ns/op	     422 B/op	      16 allocs/op
BenchmarkShard1Read8Write2-8   	  300000	      5670 ns/op	     425 B/op	      10 allocs/op
BenchmarkShard10Read8Write2-8  	  300000	      4490 ns/op	     494 B/op	      20 allocs/op
BenchmarkShard100Read8Write2-8 	 1000000	      2575 ns/op	     446 B/op	      20 allocs/op
BenchmarkShard1000Read8Write2-8	 1000000	      2658 ns/op	     476 B/op	      20 allocs/op
BenchmarkShard1Read2Write8-8   	  200000	      6520 ns/op	     586 B/op	      10 allocs/op
BenchmarkShard10Read2Write8-8  	  300000	      4607 ns/op	     668 B/op	      20 allocs/op
BenchmarkShard100Read2Write8-8 	 1000000	      2615 ns/op	     448 B/op	      20 allocs/op
BenchmarkShard1000Read2Write8-8	 1000000	      2725 ns/op	     487 B/op	      20 allocs/op
BenchmarkShard1Read8Write0-8   	 1000000	      1860 ns/op	     128 B/op	       8 allocs/op
BenchmarkShard10Read8Write0-8  	 1000000	      1987 ns/op	     257 B/op	      16 allocs/op
BenchmarkShard100Read8Write0-8 	 1000000	      1981 ns/op	     257 B/op	      16 allocs/op
BenchmarkShard1000Read8Write0-8	 1000000	      2036 ns/op	     257 B/op	      16 allocs/op
ok  	github.com/miolini/mapstore	54.959s
```