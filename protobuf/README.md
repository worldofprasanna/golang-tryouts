# Tryout with Protobuf

Golang code to create protobuf message and save the protobuf message to a file and read it back.

## Installation

1. Install protoc
2. Install protoc-go-gen using, go get -u github.com/golang/protobuf/protoc-gen-go

## To generate the protobuf go source

protoc --go_out=. msg.proto
  
## To build the binary

go build cmd/main.go

## To run the code

./main add hello there
