# mapstore
Multicore optimized in-memory data store in Go

## Benchmark

```bash
$ go test -v -cpu 16 -bench .
testing: warning: no tests to run
PASS
BenchmarkShard1Read4Write4-16         200000          6466 ns/op
BenchmarkShard10Read4Write4-16        500000          3294 ns/op
BenchmarkShard100Read4Write4-16      1000000          2052 ns/op
BenchmarkShard1000Read4Write4-16     1000000          1985 ns/op
BenchmarkShard1Read8Write2-16         300000          6458 ns/op
BenchmarkShard10Read8Write2-16        500000          3797 ns/op
BenchmarkShard100Read8Write2-16      1000000          2389 ns/op
BenchmarkShard1000Read8Write2-16     1000000          2430 ns/op
BenchmarkShard1Read2Write8-16         200000          7606 ns/op
BenchmarkShard10Read2Write8-16        500000          3531 ns/op
BenchmarkShard100Read2Write8-16      1000000          2523 ns/op
BenchmarkShard1000Read2Write8-16     1000000          2475 ns/op
ok      github.com/miolini/mapstore    36.897s
```