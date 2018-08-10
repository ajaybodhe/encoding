#!/usr/bin/env bash
export GOPATH="$(dirname "$(dirname "$(dirname "$(dirname "$(pwd)")")")")"
export PATH=$PATH:$GOPATH/bin
rm -rf models/*_easyjson*
#-leave_temps
cd models
easyjson -pkg -all
cd ..
