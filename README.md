# mapstore
Multicore optimized in-memory data store in Go

[![Build Status](https://travis-ci.org/miolini/mapstore.svg)](https://travis-ci.org/miolini/mapstore.svg) [![Go Report Card](http://goreportcard.com/badge/miolini/mapstore)](http://goreportcard.com/report/miolini/mapstore)

## Benchmark

```bash
$ go test -v -cpu 16 -bench .
=== RUN   TestShardsStat
--- PASS: TestShardsStat (12.34s)
	mapstore_test.go:35: stats: [249597 249954 249518 250931]
PASS
BenchmarkShard1Read4Write4-16   	  200000	     11549 ns/op
BenchmarkShard10Read4Write4-16  	  200000	     11488 ns/op
BenchmarkShard100Read4Write4-16 	  200000	      9974 ns/op
BenchmarkShard1000Read4Write4-16	  300000	      8340 ns/op
BenchmarkShard1Read8Write2-16   	  200000	      8683 ns/op
BenchmarkShard10Read8Write2-16  	  200000	     10528 ns/op
BenchmarkShard100Read8Write2-16 	  200000	      8799 ns/op
BenchmarkShard1000Read8Write2-16	  200000	     12656 ns/op
BenchmarkShard1Read2Write8-16   	  100000	     23423 ns/op
BenchmarkShard10Read2Write8-16  	  200000	     12483 ns/op
BenchmarkShard100Read2Write8-16 	  100000	     12010 ns/op
BenchmarkShard1000Read2Write8-16	  200000	     11880 ns/op
BenchmarkShard1Read8Write0-16   	 1000000	      2792 ns/op
BenchmarkShard10Read8Write0-16  	  300000	      3885 ns/op
BenchmarkShard100Read8Write0-16 	  500000	      4003 ns/op
BenchmarkShard1000Read8Write0-16	  500000	      4174 ns/op
ok  	github.com/miolini/mapstore	53.613s
```