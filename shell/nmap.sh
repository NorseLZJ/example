#!/usr/bin/env bash

ip=$1

split_str="*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-"

if [ -z "$ip" ]; then
    echo "param err"
    exit
fi

cat /dev/null >nmap_report.txt
lst=$(nmap -sP ${ip}'/24' | grep 'Nmap scan report for' | awk '{print $5}')

for i in ${lst}; do
    {
        report="nmap_report.txt"
        $(nmap -F -T5 --version-light --top-ports 300 ${i} >>${report})
        echo ${split_str} >>nmap_report.txt
    } &
done
wait
echo "check done."
