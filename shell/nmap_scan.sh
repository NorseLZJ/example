#!/bin/env bash

#ifconfig | grep 'inet' | grep -E -v 'inet6|127|172'

ip=$(traceroute www.baidu.com -n -m1 | grep '192' | awk '{print $2}')
echo $ip
nmap -sP $ip/24 -T5  
