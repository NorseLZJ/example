#!/bin/env bash

lst=(
    'lzj'
    'norselzj'
)

cat /dev/null >upwd.txt
for i in ${lst[@]}; do
    #echo ${i}
    pwd=$(echo ${i} | md5sum)
    echo -e ${i}":"${pwd} | tee -a upwd.txt
done
