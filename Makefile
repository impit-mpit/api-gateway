API_PROTO_FILES=$(shell find api-proto/src/proto -name *.proto)
.PHONY: generate
generate:
	protoc --proto_path=./api-proto/src \
		   --proto_path=./api-proto/src/third_party \
 	       --grpc-gateway_out=paths=source_relative:./gen/go \
		   --grpc-gateway_opt=paths=source_relative \
		   --grpc-gateway_opt generate_unbound_methods=true \
	       $(API_PROTO_FILES)