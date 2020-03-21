#!/bin/bash

if [ "$1" = 'windows' ]; then
export GOOS=linux 
export GOARCH=amd64
go build -ldflags='-w -s' -o $GOPATH/src/github.com/NorseLZJ/program/fileServer/fileServer.exe 
else 
go build -ldflags='-w -s' -o $GOPATH/src/github.com/NorseLZJ/program/fileServer/fileServer 
fi
