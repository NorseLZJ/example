#!/bin/bash

if [ "$1" = 'init' ]; 
then
docker run \
    --name nsq1 \
    -p 4150:4150 \
    -p 4151:4151 \
    -p 4160:4160 \
    -p 4161:4161 \
    --restart always \
    -d s_nsqd  
elif [ "$1" = 'in' ]; 
then
docker exec -it nsq1 /bin/sh
else
echo "no match option!!!"
fi
