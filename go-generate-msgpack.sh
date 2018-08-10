#!/usr/bin/env bash
export GOPATH="$(dirname "$(dirname "$(dirname "$(dirname "$(pwd)")")")")"
export PATH=$PATH:$GOPATH/bin
rm -rf models/*data_gen*
#-leave_temps
cd models
go generate
cd ..