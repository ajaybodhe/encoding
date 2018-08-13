#!/usr/bin/env bash
export GOPATH="$(dirname "$(dirname "$(dirname "$(dirname "$(pwd)")")")")"
export PATH=$PATH:$GOPATH/bin
rm -rf models/gogModels/*pb.go*
#-leave_temps
cd models/gogoModels
protoc --gofast_out=. gogo.proto
cd ../..
