all: go_proto
	
go_proto:
	protoc --go_out=. --go-grpc_out=require_unimplemented_servers=false:. auth.proto
	protoc --go_out=. --go-grpc_out=require_unimplemented_servers=false:. course.proto
	protoc --go_out=. --go-grpc_out=require_unimplemented_servers=false:. exam.proto