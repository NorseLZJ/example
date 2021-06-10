#!/bin/env bash

cmd=$1
if [[ $cmd != "" &&  $cmd == "scan" ]];then
    nmap -p 22 -oX ssh.xml 192.168.1.1-251
fi

grep -n -B 4 'open' ssh.xml | grep 'address addr=' | grep -v 'mac' | awk '{print $2}'
