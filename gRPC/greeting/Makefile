PROTO_OUT ?= pkg/pb

DIR ?= $(shell pwd)

PROTOC_CMD ?= docker run --rm -v $(DIR):$(DIR) -w $(DIR) thethingsindustries/protoc

.PHONY: compiler-proto-go
compile-proto-go:
	@rm -rf $(PROTO_OUT)
	@mkdir -p $(PROTO_OUT)
	@find proto -name "*.proto" -type f -exec $(PROTOC_CMD) -Iproto/ \
		--go_opt=paths=source_relative \
		--go_out=plugins=grpc:$(PROTO_OUT) {} \;