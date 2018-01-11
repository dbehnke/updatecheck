#!/bin/bash
set -e -x
export CURRENTDIR=$PWD
export GOPATH=/tmp/updatecheck/go
export PROJECTPATH=$GOPATH/src/gitlab.com/dbehnke74/updatecheck
mkdir -p $PROJECTPATH || true
cp -R * $PROJECTPATH
go build cmd/updatecheck/main.go
./main results
diff -c1 checkresults.json results.json || \
  diff -C 1 checkresults.json results.json

