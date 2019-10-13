build:
	protoc -I/usr/local/include -I./proto \
	-I$$GOPATH/src -I$$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--go_out=plugins=grpc:../engine/lib/structs \
	proto/*.proto
gateway: build
	protoc -I/usr/local/include -I./proto \
	-I$$GOPATH/src -I$$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--grpc-gateway_out=logtostderr=true:../engine/lib/structs \
	proto/*.proto
proto: gateway
	for file in `find ../engine/lib/structs -type f -name "*.pb.go"`; \
	 do \
	  protoc-go-inject-tag -input="$$file"; \
	done;
