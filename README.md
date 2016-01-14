# mapstore
Multicore optimized in-memory data store in Go

[![Build Status](https://travis-ci.org/miolini/mapstore.svg)](https://travis-ci.org/miolini/mapstore.svg) [![Go Report Card](http://goreportcard.com/badge/miolini/mapstore)](http://goreportcard.com/report/miolini/mapstore)

## Benchmark

```bash
$ go test -v -cpu 16 -bench .
=== RUN   TestShardsStat
--- PASS: TestShardsStat (4.93s)
	mapstore_test.go:35: stats: [249597 249954 249518 250931]
PASS
BenchmarkShard1Read4Write4-16   	  300000	      5875 ns/op
BenchmarkShard10Read4Write4-16  	  500000	      3381 ns/op
BenchmarkShard100Read4Write4-16 	 1000000	      2031 ns/op
BenchmarkShard1000Read4Write4-16	 1000000	      2008 ns/op
BenchmarkShard1Read8Write2-16   	  300000	      6440 ns/op
BenchmarkShard10Read8Write2-16  	  500000	      3824 ns/op
BenchmarkShard100Read8Write2-16 	 1000000	      2403 ns/op
BenchmarkShard1000Read8Write2-16	 1000000	      2466 ns/op
BenchmarkShard1Read2Write8-16   	  200000	      7583 ns/op
BenchmarkShard10Read2Write8-16  	  300000	      3591 ns/op
BenchmarkShard100Read2Write8-16 	 1000000	      2508 ns/op
BenchmarkShard1000Read2Write8-16	 1000000	      2529 ns/op
BenchmarkShard1Read8Write0-16   	 1000000	      2031 ns/op
BenchmarkShard10Read8Write0-16  	 1000000	      1827 ns/op
BenchmarkShard100Read8Write0-16 	 1000000	      1798 ns/op
BenchmarkShard1000Read8Write0-16	 1000000	      1795 ns/op
ok  	github.com/miolini/mapstore	53.613s
```