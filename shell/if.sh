#!/bin/env bash

#if [ 'hello' = 'world' ];then
#        echo 'hello == world' 
#else
#        echo 'hello != world' 
#fi

val1=$1
val2=$2
val3=$3
if [ -z ${val1} || -z ${val2} || -z ${val3} ];then
    echo 'check your param'
fi

echo 'hello'
echo 'hello'
echo 'hello'
echo 'hello'
