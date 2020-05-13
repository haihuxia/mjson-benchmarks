## mjson-benchmarks

Benchmarks for [mjson]().

*These benchmarks were run on a MacBook Pro 15" Mid 2014 2.5 GHz Intel Core i7 using Go 1.14.*

### Usage

```
go test -bench=. -count=3
```

* v1.0.2 vs v1.0.0 

BenchmarkMappingString-8：22%、18%、20%<br/>BenchmarkMappingSpec-8：v1.0.0 has several bugs<br/>BenchmarkMappingMap-8：27%、28%、26%

```
goos: darwin
goarch: amd64
pkg: mjson-benchmarks
BenchmarkMappingString-8   	  968082	      1093 ns/op	    1792 B/op	       2 allocs/op
BenchmarkMappingString-8   	 1000000	      1078 ns/op	    1792 B/op	       2 allocs/op
BenchmarkMappingString-8   	 1000000	      1082 ns/op	    1792 B/op	       2 allocs/op
BenchmarkMappingSpec-8     	  174841	      6654 ns/op	   10752 B/op	      12 allocs/op
BenchmarkMappingSpec-8     	  177376	      6644 ns/op	   10752 B/op	      12 allocs/op
BenchmarkMappingSpec-8     	  176644	      6664 ns/op	   10752 B/op	      12 allocs/op
BenchmarkMappingMap-8      	  352795	      3325 ns/op	    5376 B/op	       6 allocs/op
BenchmarkMappingMap-8      	  353355	      3330 ns/op	    5376 B/op	       6 allocs/op
BenchmarkMappingMap-8      	  350379	      3329 ns/op	    5376 B/op	       6 allocs/op
PASS
ok  	mjson-benchmarks	10.626s
```

* v1.0.0

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

