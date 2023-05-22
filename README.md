# Golang Plugin Benchmark

This is a benchmark of various Go plugin implementations (native, wasm, rpc, ...).

## Results

```
goos: darwin
goarch: amd64
pkg: github.com/euskadi31/go-plugin-benchmark
cpu: Intel(R) Core(TM) i7-6820HQ CPU @ 2.70GHz
BenchmarkGoNativeFuncInline-8     	1000000000	    0.3477 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoNativeFuncNoInline-8   	 755171646	     1.584 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoNativePlugin-8         	 757251984	     1.579 ns/op	       0 B/op	       0 allocs/op
BenchmarkHashicorpPlugin-8        	     20880	     60441 ns/op	     584 B/op	      15 allocs/op
BenchmarkWasmerRustSum-8          	    328557	      3989 ns/op	     208 B/op	      12 allocs/op
BenchmarkWasmerGoSum-8            	    326815	      3555 ns/op	     208 B/op	      12 allocs/op
BenchmarkWazeroRustSum-8          	  13622338	     81.71 ns/op	      24 B/op	       2 allocs/op
BenchmarkWazeroGoSum-8            	  14121835	     82.20 ns/op	      24 B/op	       2 allocs/op
BenchmarkYaegi-8                  	   1261688	     920.8 ns/op	     528 B/op	      14 allocs/op
PASS
coverage: [no statements]
ok  	github.com/euskadi31/go-plugin-benchmark	12.605s
```
