#!/bin/bash

# 使用代理执行某个命令

echo "export and ... "

export http_proxy=http://192.168.1.120:7890
export https_proxy=http://192.168.1.120:7890

$*
