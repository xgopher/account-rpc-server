# account-rpc-server
用户中心 - RPC服务端


## 安装

```
export GOPROXY=https://goproxy.io
go build ./...
./app
```


## Generate gRPC code

```
protoc -I helloworld/ helloworld/helloworld.proto --go_out=plugins=grpc:helloworld
```


## 运行服务/客户端

```
go run main.go
go run client/main.go
```
