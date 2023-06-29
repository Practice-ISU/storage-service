go_gen:
	protoc --go_out=../../pkg/grpc/file --go_opt=paths=source_relative file.proto
	protoc --go-grpc_out=../../pkg/grpc/file --go-grpc_opt=paths=source_relative file.proto