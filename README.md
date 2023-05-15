# Golang Plugin Benchmark

This is a benchmark of various Go plugin implementations (native, wasm).

## Results

```
goos: darwin
goarch: amd64
pkg: github.com/euskadi31/go-plugin-benchmark
cpu: Intel(R) Core(TM) i7-6820HQ CPU @ 2.70GHz
BenchmarkGoPlugin-8             610472167                2.106 ns/op           0 B/op          0 allocs/op
BenchmarkWasmerRustSum-8          325718                 3560 ns/op          208 B/op         12 allocs/op
BenchmarkWasmerGoSum-8            339010                 3545 ns/op          208 B/op         12 allocs/op
BenchmarkWazeroRustSum-8        13045299                79.12 ns/op           24 B/op          2 allocs/op
BenchmarkWazeroGoSum-8          12553725                82.37 ns/op           24 B/op          2 allocs/op
PASS
ok      github.com/euskadi31/go-plugin-benchmark        7.562s
```
