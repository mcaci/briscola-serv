#!/bin/sh
set -ev
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
protoc briscola/pb/*.proto --go_out=plugins=grpc:.
echo