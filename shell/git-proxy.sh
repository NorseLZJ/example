#!/bin/env bash

git config --global http.proxy socks5://127.0.0.1:1080
git config --global https.proxy socks5://127.0.0.1:1080

git config --global https.proxy http://127.0.0.1:1080
git config --global https.proxy https://127.0.0.1:1080
