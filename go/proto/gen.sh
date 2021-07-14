#!/bin/env bash

wd=$(dirname $(readlink -f "$0"))

SRC_DIR=$wd
DST_DIR=$wd

protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/login.proto
protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/game.proto
