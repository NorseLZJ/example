#!/bin/env bash

wd=$(dirname $(readlink -f "$0"))

SRC_DIR=$wd
DST_DIR=$wd

protoc -I=$SRC_DIR --gofast_out=$DST_DIR $SRC_DIR/wsgate.proto

cp -rf pb ../wsgate/
cp -rf pb ../wscli/
