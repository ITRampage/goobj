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
%  go version
go version go1.18 darwin/amd64

% go test -bench=.
goos: darwin
goarch: amd64
pkg: github.com/ITRampage/goobj/inc
cpu: Intel(R) Core(TM) i7-4770HQ CPU @ 2.20GHz
BenchmarkObj-8                   2  514196514 ns/op
BenchmarkInc-8                3918     301594 ns/op
BenchmarkCIncNoLock-8         1358     848450 ns/op
BenchmarkCInc-8                 10  105194498 ns/op
BenchmarkCIncLockAll-8         534    2244140 ns/op
PASS
ok   github.com/ITRampage/goobj/inc 6.715s

```
```
% go version
go version go1.17.3 darwin/amd64

% go test -bench=.
goos: darwin
goarch: amd64
pkg: github.com/ITRampage/goobj/inc
cpu: Intel(R) Core(TM) i7-8850H CPU @ 2.60GHz
BenchmarkObj-12                    3  364824308 ns/op
BenchmarkInc-12                 2463     486396 ns/op
BenchmarkCIncNoLock-12          1296     871513 ns/op
BenchmarkCInc-12                  22   50133538 ns/op
BenchmarkCIncLockAll-12          757    1527242 ns/op
PASS
ok   github.com/ITRampage/goobj/inc 8.182s
```
```
% go version
go version go1.19 darwin/amd64

% go test -bench=.
goos: darwin
goarch: amd64
pkg: github.com/ITRampage/goobj/inc
cpu: Intel(R) Core(TM) i7-8850H CPU @ 2.60GHz
BenchmarkObj-12                    3  360773165 ns/op
BenchmarkInc-12                 5036     241384 ns/op
BenchmarkCIncNoLock-12          1275     848847 ns/op
BenchmarkCInc-12                  22   50639908 ns/op
BenchmarkCIncLockAll-12          753    1532934 ns/op
PASS
ok   github.com/ITRampage/goobj/inc 8.163s
```
