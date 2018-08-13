# encoding
testing and benchmarking encoders in golang : https://github.com/alecthomas/go_serialization_benchmarks

go get -v gopkg.in/src-d/proteus.v1/...

go get -u github.com/andyleap/gencode 

go get -u github.com/pascaldekloe/colfer/cmd/colf

go get github.com/linkedin/goavro.v2 

go get  "github.com/davecgh/go-xdr/xdr2"

go get  github.com/tinylib/msgp 

go get github.com/gogo/protobuf/proto   
                                         
go get github.com/gogo/protobuf/jsonpb

go get github.com/gogo/protobuf/protoc-gen-gogo

go get github.com/gogo/protobuf/gogoproto

go get github.com/gogo/protobuf/protoc-gen-gofast

go get -u github.com/mailru/easyjson/...

git clone

cd testing

go test -bench=. -benchmem


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



