#!/usr/bin/env bash
export GOPATH="$(dirname "$(dirname "$(dirname "$(dirname "$(pwd)")")")")"
export PATH=$PATH:$GOPATH/bin
rm -rf models/gencodeModels/*gen.go*
#-leave_temps
cd models/gencodeModels
gencode go -schema gencode.schema -package gencodeModels
cd ../..
