#!/bin/bash

if [ "$1" = 'windows' ]; then
export GOOS=windows
export GOARCH=amd64
export GOHOSTARCH=amd64
export GOHOSTOS=windows
export GOEXE=.exe
go build -ldflags='-w -s' -o $GOPATH/src/github.com/NorseLZJ/program/fileServer/fileServer.exe 
else 
go build -ldflags='-w -s' -o $GOPATH/src/github.com/NorseLZJ/program/fileServer/fileServer 
fi
