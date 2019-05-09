#!/bin/bash

## 查找.proto文件的目录 + path/to/file.proto + out输出目录
protoc -I protos user/user.proto --go_out=plugins=grpc:protos