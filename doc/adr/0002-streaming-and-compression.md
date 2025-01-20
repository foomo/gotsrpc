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
BenchmarkBufferedClient/none
BenchmarkBufferedClient/none/0
BenchmarkBufferedClient/none/0-10       	   22604	     49513 ns/op	   28943 B/op	     119 allocs/op
BenchmarkBufferedClient/none/1
BenchmarkBufferedClient/none/1-10       	   23991	     55106 ns/op	   28862 B/op	     119 allocs/op
BenchmarkBufferedClient/none/2
BenchmarkBufferedClient/none/2-10       	   22976	     50808 ns/op	   28946 B/op	     119 allocs/op
BenchmarkBufferedClient/gzip
BenchmarkBufferedClient/gzip/0
BenchmarkBufferedClient/gzip/0-10       	    9292	    124257 ns/op	   15796 B/op	     117 allocs/op
BenchmarkBufferedClient/gzip/1
BenchmarkBufferedClient/gzip/1-10       	   10520	    112287 ns/op	   15601 B/op	     117 allocs/op
BenchmarkBufferedClient/gzip/2
BenchmarkBufferedClient/gzip/2-10       	    9838	    125777 ns/op	   15751 B/op	     117 allocs/op
BenchmarkBufferedClient/snappy
BenchmarkBufferedClient/snappy/0
BenchmarkBufferedClient/snappy/0-10     	   21604	     62413 ns/op	   15189 B/op	     114 allocs/op
BenchmarkBufferedClient/snappy/1
BenchmarkBufferedClient/snappy/1-10     	   21208	     54509 ns/op	   15242 B/op	     114 allocs/op
BenchmarkBufferedClient/snappy/2
BenchmarkBufferedClient/snappy/2-10     	   24153	     51172 ns/op	   15253 B/op	     114 allocs/op
```
