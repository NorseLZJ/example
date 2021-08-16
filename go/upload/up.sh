#!/usr/bin/env bash

#curl -F "file=@/d/share/Downloads/hyper-Setup-3.0.2.exe" -F "key=17c813eefe8c34998867a8c92e67757a164659b247dba9da7bdbba664f52f629" -X POST http://192.168.1.116:6500/upload
curl -F "file=@/d/share/Downloads/hyper-Setup-3.0.2.exe" -F "key=123456" -X POST http://192.168.1.116:6500/upload
