.PHONY: protos

protos:
	protoc --proto_path=protos --go_out=protos/chat --go_opt=paths=source_relative protos/chat.proto
	protoc --proto_path=protos --go_out=protos/chat --go_opt=paths=source_relative --go-grpc_out=protos/chat protos/chat.proto