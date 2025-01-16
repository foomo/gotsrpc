# 2. streaming and compression

Date: 2025-01-16

## Status

Accepted

## Context

We need to allow streaming and compression of the data.

## Decision

We will use the `io.Pipe` interface to allow streaming of the data, and pipe the data through the compression algorithm.
For 

## Consequences

Benchmarking shows that the compression is not noticeable for snappy, but the streaming is in comparison to the old method.

```text
goos: darwin
goarch: arm64
pkg: github.com/foomo/gotsrpc/v2
cpu: Apple M1 Max
BenchmarkBufferedClient
BenchmarkBufferedClient/deprecated
BenchmarkBufferedClient/deprecated/0
BenchmarkBufferedClient/deprecated/0-10       	   20712	     54428 ns/op	   26135 B/op	     108 allocs/op
BenchmarkBufferedClient/deprecated/1
BenchmarkBufferedClient/deprecated/1-10       	   22180	     52320 ns/op	   26170 B/op	     108 allocs/op
BenchmarkBufferedClient/deprecated/2
BenchmarkBufferedClient/deprecated/2-10       	   22329	     49323 ns/op	   26236 B/op	     108 allocs/op
BenchmarkBufferedClient/deprecated/3
BenchmarkBufferedClient/deprecated/3-10       	   24580	     48177 ns/op	   26138 B/op	     108 allocs/op
BenchmarkBufferedClient/deprecated/4
BenchmarkBufferedClient/deprecated/4-10       	   24999	     57772 ns/op	   26154 B/op	     108 allocs/op
BenchmarkBufferedClient/snappy
BenchmarkBufferedClient/snappy/0
BenchmarkBufferedClient/snappy/0-10           	   16392	     69553 ns/op	   15282 B/op	     114 allocs/op
BenchmarkBufferedClient/snappy/1
BenchmarkBufferedClient/snappy/1-10           	   17702	     72923 ns/op	   14944 B/op	     114 allocs/op
BenchmarkBufferedClient/snappy/2
BenchmarkBufferedClient/snappy/2-10           	   17932	     67446 ns/op	   14819 B/op	     114 allocs/op
BenchmarkBufferedClient/snappy/3
BenchmarkBufferedClient/snappy/3-10           	   16640	     69216 ns/op	   15155 B/op	     114 allocs/op
BenchmarkBufferedClient/snappy/4
BenchmarkBufferedClient/snappy/4-10           	   16767	     66247 ns/op	   15010 B/op	     114 allocs/op
BenchmarkBufferedClient/none
BenchmarkBufferedClient/none/0
BenchmarkBufferedClient/none/0-10             	   17706	     68516 ns/op	   12280 B/op	     112 allocs/op
BenchmarkBufferedClient/none/1
BenchmarkBufferedClient/none/1-10             	   17593	     68580 ns/op	   12308 B/op	     112 allocs/op
BenchmarkBufferedClient/none/2
BenchmarkBufferedClient/none/2-10             	   17292	     67673 ns/op	   12208 B/op	     112 allocs/op
BenchmarkBufferedClient/none/3
BenchmarkBufferedClient/none/3-10             	   17086	     71715 ns/op	   12285 B/op	     112 allocs/op
BenchmarkBufferedClient/none/4
BenchmarkBufferedClient/none/4-10             	   17067	     68955 ns/op	   12295 B/op	     112 allocs/op
BenchmarkBufferedClient/gzip
BenchmarkBufferedClient/gzip/0
BenchmarkBufferedClient/gzip/0-10             	    7190	    153284 ns/op	   22024 B/op	     113 allocs/op
BenchmarkBufferedClient/gzip/1
BenchmarkBufferedClient/gzip/1-10             	    6808	    158344 ns/op	   20757 B/op	     113 allocs/op
BenchmarkBufferedClient/gzip/2
BenchmarkBufferedClient/gzip/2-10             	    6889	    156492 ns/op	   19680 B/op	     113 allocs/op
BenchmarkBufferedClient/gzip/3
BenchmarkBufferedClient/gzip/3-10             	    6927	    148146 ns/op	   18912 B/op	     113 allocs/op
BenchmarkBufferedClient/gzip/4
BenchmarkBufferedClient/gzip/4-10             	    8340	    146697 ns/op	   20207 B/op	     113 allocs/op
```
