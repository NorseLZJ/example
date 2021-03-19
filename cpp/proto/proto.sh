#!/bin/bash

cmd=$1

echo "gen proto ....."

workdir=$(
    cd $(dirname $0)
    pwd
)

echo ${workdir}

gen_proto() {
    local pub_dir=""
    local SRC_DIR=${workdir}"/"
    local DST_DIR=${workdir}"/"
    protoc -I=${SRC_DIR} --cpp_out=${DST_DIR} ${SRC_DIR}netcmd.proto
}

case ${cmd} in
"pb")
    echo "gen proto to cc & h ..."
    gen_proto
    ;;
*)
    echo "no option"
    ;;
esac