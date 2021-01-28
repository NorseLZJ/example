#!/bin/env bash

lst=(
    'gubo'
)

cat /dev/null >upwd.txt
for i in ${lst[@]}; do

    pwd=$(echo ${i} | md5sum)
    echo -e ${i}":"${pwd} | tee -a upwd.txt
done
