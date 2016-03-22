#!/usr/bin/env bash

export PATH="/home/rogeralsing/go/bin:$PATH"

go get -u github.com/gengo/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/gengo/grpc-gateway/protoc-gen-swagger
go get -u github.com/golang/protobuf/protoc-gen-go

echo "gRPC"
protoc -I/usr/local/include -I. \
 -I$GOPATH/src \
 -I$GOPATH/src/github.com/gengo/grpc-gateway/third_party/googleapis \
 --go_out=Mgoogle/api/annotations.proto=github.com/gengo/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:. \
 *.proto

echo "REST API"
protoc -I/usr/local/include -I. \
 -I$GOPATH/src \
 -I$GOPATH/src/github.com/gengo/grpc-gateway/third_party/googleapis \
 --grpc-gateway_out=logtostderr=true:. \
 *.proto

echo "Swagger"
protoc -I/usr/local/include -I. \
 -I$GOPATH/src \
 -I$GOPATH/src/github.com/gengo/grpc-gateway/third_party/googleapis \
 --swagger_out=logtostderr=true:. \
 *.proto