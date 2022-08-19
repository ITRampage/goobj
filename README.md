# The example of objects implementations in Golang

## How to run example?

```
% cd <project path>
% go run target/main.go
```

## How to run benchmarks?

```
% cd <project path>/inc
% go test -bench=.
```

My results are:

```
% go test -bench=.
goos: darwin
goarch: amd64
pkg: github.com/ITRampage/goobj/inc
cpu: Intel(R) Core(TM) i7-4770HQ CPU @ 2.20GHz
BenchmarkObj-8                   2  525474538 ns/op
BenchmarkInc-8                3926     306248 ns/op
BenchmarkCIncNoLock-8         1358     891848 ns/op
BenchmarkCInc-8                 10  105827823 ns/op
BenchmarkCIncLockAll-8         532    2243374 ns/op

```
