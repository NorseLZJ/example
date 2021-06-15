#!/bin/env bash

cmd=$1
if [[ $cmd != "" ]];then
    nmap -p 8080 -oX jenkins.xml 192.168.1.1-251
fi

grep -n -B 4 'open' jenkins.xml | grep 'address addr=' | grep -v 'mac' | awk '{print $2}'
