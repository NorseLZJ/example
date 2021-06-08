#!/bin/bash

echo "--------- IMAEGS ------------"
docker images

echo "--------- CONTAINER ------------"
docker ps -qa

docker rm $(docker ps -aq)

ret=$(docker images | grep none)
if [[ $ret != "" ]];then
    echo $ret
fi

ret=""
ret=$(docker images | grep 'test/app')

if [[ $ret != "" ]];then
    id=$(echo $ret | awk '{print $3}')
    echo $id
    docker rmi $id
fi

docker build -t test/app . 
