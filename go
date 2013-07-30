#!/bin/bash

export GOROOT=$HOME/milo/build/go
export PATH=$PATH:$GOROOT/bin

export GOPATH=$HOME/milo/web/gowf:$GOROOT/libs

go "$@"
