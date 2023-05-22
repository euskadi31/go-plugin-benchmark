# Golang Plugin Benchmark

This is a benchmark of various Go plugin implementations (native, wasm).

## Results

```
goos: darwin
goarch: amd64
pkg: github.com/euskadi31/go-plugin-benchmark
cpu: Intel(R) Core(TM) i7-6820HQ CPU @ 2.70GHz
BenchmarkGoNativePlugin-8      754253935	     1.628 ns/op	       0 B/op	       0 allocs/op
BenchmarkHashicorpPlugin-8   	   19508	     64092 ns/op	     584 B/op	      15 allocs/op
BenchmarkWasmerRustSum-8     	  332524	      3808 ns/op	     208 B/op	      12 allocs/op
BenchmarkWasmerGoSum-8       	  293564	      3584 ns/op	     208 B/op	      12 allocs/op
BenchmarkWazeroRustSum-8     	12958280	     86.12 ns/op	      24 B/op	       2 allocs/op
BenchmarkWazeroGoSum-8       	11966158	     87.66 ns/op	      24 B/op	       2 allocs/op
PASS
coverage: [no statements]
ok  	github.com/euskadi31/go-plugin-benchmark	8.531s
```
