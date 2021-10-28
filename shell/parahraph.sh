#!/usr/bin/env bash

#扫网段
for ((i = 1;i < 255;i++));
do 
	ip="192.168."$i".1"
	report=$(nmap -T5 -sP $ip | grep 'Host is up')
	if [[ $report != "" ]];then
		printf "%s -> %s\n" $i $ip
	fi
done
