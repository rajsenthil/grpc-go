#!/bin/bash
echo "Generating protoc..."
# protoc ./greet/greetpb/greet.proto --plugin=$(go env GOPATH)/bin/protoc-gen-go-grpc --go-grpc_out=./greet/greetpb/pb
# protoc ./greet/greetpb/greet.proto --plugin=$(go env GOPATH)/bin/protoc-gen-go --go-grpc_out=./greet/greetpb
# protoc --go_out=. --go_opt=paths=./greet/greetpb/pb \
#     --go-grpc_out=. --go-grpc_opt=paths=./greet/greetpb/pb \
#     ./greet/greetpb/greet.proto
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./greet/greetpb/greet.proto
echo "Done."
