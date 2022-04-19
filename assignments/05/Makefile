.PHONY: again

# generate gRPC code
compile:
	protoc -I=. --go_out=. --go-grpc_out=./server --go-grpc_opt=paths=source_relative server.proto
