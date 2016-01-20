# mapstore
Multicore optimized in-memory data store in Go

[![Build Status](https://travis-ci.org/miolini/mapstore.svg)](https://travis-ci.org/miolini/mapstore.svg) [![Go Report Card](http://goreportcard.com/badge/miolini/mapstore)](http://goreportcard.com/report/miolini/mapstore)

## Benchmark
go test -v -cpu 24 -bench .
```
$ benchcmp -best miolini.txt serjvanilla.txt
benchmark                            old ns/op     new ns/op     delta
BenchmarkShard1Read4Write4-24        9612          7981          -16.97%
BenchmarkShard10Read4Write4-24       6121          6010          -1.81%
BenchmarkShard100Read4Write4-24      3921          3943          +0.56%
BenchmarkShard1000Read4Write4-24     3272          3274          +0.06%
BenchmarkShard1Read8Write2-24        11128         9963          -10.47%
BenchmarkShard10Read8Write2-24       6408          6267          -2.20%
BenchmarkShard100Read8Write2-24      3755          3852          +2.58%
BenchmarkShard1000Read8Write2-24     3647          3736          +2.44%
BenchmarkShard1Read2Write8-24        11313         9817          -13.22%
BenchmarkShard10Read2Write8-24       6822          6992          +2.49%
BenchmarkShard100Read2Write8-24      5008          5134          +2.52%
BenchmarkShard1000Read2Write8-24     4091          4145          +1.32%
BenchmarkShard1Read8Write0-24        2302          1677          -27.15%
BenchmarkShard10Read8Write0-24       2281          2360          +3.46%
BenchmarkShard100Read8Write0-24      2278          2350          +3.16%
BenchmarkShard1000Read8Write0-24     2296          2389          +4.05%
```
