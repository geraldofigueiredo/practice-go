.PHONY: protos

protos:
	protoc --proto_path=protos --go_out=protos/menu --go_opt=paths=source_relative \
    --go-grpc_out=protos/menu --go-grpc_opt=paths=source_relative \
    protos/menu.proto