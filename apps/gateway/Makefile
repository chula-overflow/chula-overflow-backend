all: install build

install: docs dependency

dependency: protoc
	go mod download

docs: protoc
	swag init -o docs

protoc:
	protoc --proto_path=../proto/ --go_out=. --go-grpc_out=require_unimplemented_servers=false:. ../proto/auth.proto
	protoc --proto_path=../proto/ --go_out=. --go-grpc_out=require_unimplemented_servers=false:. ../proto/course.proto
	protoc --proto_path=../proto/ --go_out=. --go-grpc_out=require_unimplemented_servers=false:. ../proto/exam.proto
	protoc --proto_path=../proto/ --go_out=. --go-grpc_out=require_unimplemented_servers=false:. ../proto/thread.proto

build: