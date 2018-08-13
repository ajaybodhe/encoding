#!/usr/bin/env bash
export GOPATH="$(dirname "$(dirname "$(dirname "$(dirname "$(pwd)")")")")"
export PATH=$PATH:$GOPATH/bin
rm -fr models/generated.pb.go models/proteusGeneratedModels
mkdir models/proteusGeneratedModels
proteus   -f models/proteusGeneratedModels -p github.com/ajaybodhe/encoding/models --verbose
