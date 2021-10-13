#!/bin/env bash

cwd=$(readlink -f "$(dirname "$0")")
echo $cwd

tagFile=$cwd"/main.log"

if [[ ! -f tagFile ]];
then
    touch $tagFile
fi

ps -ef | wc -l >> $tagFile
