## mjson-benchmarks

Benchmarks for [mjson]().

*These benchmarks were run on a MacBook Pro 15" Mid 2014 2.5 GHz Intel Core i7 using Go 1.14.*

### Usage

```
go test -bench=. -count=3
```

```
goos: darwin
goarch: amd64
pkg: mjson-benchmarks
BenchmarkMappingString-8   	  751236	      1436 ns/op	    1792 B/op	       2 allocs/op
BenchmarkMappingString-8   	  816278	      1442 ns/op	    1792 B/op	       2 allocs/op
BenchmarkMappingString-8   	  802910	      1435 ns/op	    1792 B/op	       2 allocs/op
BenchmarkMappingSpec-8     	  271306	      4242 ns/op	    7168 B/op	       8 allocs/op
BenchmarkMappingSpec-8     	  274178	      4243 ns/op	    7168 B/op	       8 allocs/op
BenchmarkMappingSpec-8     	  272310	      4264 ns/op	    7168 B/op	       8 allocs/op
BenchmarkMappingMap-8      	  257775	      4475 ns/op	    5376 B/op	       6 allocs/op
BenchmarkMappingMap-8      	  255393	      4470 ns/op	    5376 B/op	       6 allocs/op
BenchmarkMappingMap-8      	  259144	      4475 ns/op	    5376 B/op	       6 allocs/op
PASS
ok  	mjson-benchmarks	10.800s
```

