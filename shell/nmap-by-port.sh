#!/bin/bash

port=$1
if [[ $port -eq "" ]];then
    echo "Usage"
    echo "./nmap-by-port.sh port[1~65535]"
    exit 1
fi

getway=$(ip route | grep default | awk '{print $3}')
echo "getway .. "$getway
nmap -sP $getway/24 -oG ip.xml > /dev/null
cat ip.xml | grep 192 | awk '{print $2}'  > ip.txt

lst=$(cat ip.txt)

for i in ${lst[@]}; do
    val=${#i}
    if [[ $val -gt 5 ]];then
        ret=$(nmap -p $port $i | grep open)
        if [[ $ret != "" ]];then
            echo $i
        fi
    fi
done

