#!/bin/bash
echo "Generating calc protoc..."
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./pb/FibonocciSeries.proto
echo "Done."
