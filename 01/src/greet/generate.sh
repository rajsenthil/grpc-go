#!/bin/bash
echo "Generating protoc..."
protoc ./greetpb/greet.proto --plugin=$(go env GOPATH)/bin/protoc-gen-go --go-grpc_out=./greetpb/pb
echo "Done."
