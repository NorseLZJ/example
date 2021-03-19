#!/bin/bash 

if [ ! -f "$PWD/main" ];then
    make main;
else
    rm -f "$PWD/main";
    make main;
fi