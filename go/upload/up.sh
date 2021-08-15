#!/usr/bin/env bash

curl -F "file=@/home/rs/Downloads/gmys_server.zip" -F "sign=123456" -X POST http://localhost:6500/upload
  #-H "Content-Type: multipart/form-data"
