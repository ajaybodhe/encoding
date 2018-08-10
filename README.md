# encoding
testing and benchmarking encoders in golang

git clone

cd testing

go test -bench=.


➜  testing go test -bench=.

goos: linux

goarch: amd64

pkg: github.com/ajaybodhe/encoding/testing

BenchmarkJson-4      	  200000	      6960 ns/op

BenchmarkMsgPack-4   	  200000	      7562 ns/op

BenchmarkXDR-4       	  300000	      5292 ns/op

BenchmarkAvro-4      	  300000	      4719 ns/op

BenchmarkGob-4       	   20000	     61458 ns/op

PASS

ok  	github.com/ajaybodhe/encoding/testing	8.046s



