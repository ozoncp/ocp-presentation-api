.PHONY: build
build: vendor-proto .generate .build

PHONY: .generate
.generate:
		mkdir -p pkg/ocp-presentation-api
		protoc -I vendor.protogen \
				--go_out=pkg/ocp-presentation-api --go_opt=paths=import \
				--go-grpc_out=pkg/ocp-presentation-api --go-grpc_opt=paths=import \
				--grpc-gateway_out=pkg/ocp-presentation-api \
				--grpc-gateway_opt=logtostderr=true \
				--grpc-gateway_opt=paths=import \
				--validate_out lang=go:pkg/ocp-presentation-api \
				--swagger_out=allow_merge=true,merge_file_name=api:. \
				api/ocp-presentation-api/ocp-presentation-api.proto
		mv pkg/ocp-presentation-api/gihtub.com/ozoncp/ocp-presentation-api/pkg/ocp-presentation-api/* pkg/ocp-presentation-api/
		rm -rf pkg/ocp-presentation-api/gihtub.com
		mkdir -p cmd/ocp-presentation-api

PHONY: .build
.build:
		go build -o bin/ocp-presentation-api cmd/ocp-presentation-api/main.go

PHONY: install
install: build .install

PHONY: .install
install:
		go install cmd/grpc-server/main.go

PHONY: vendor-proto
vendor-proto: .vendor-proto

PHONY: .vendor-proto
.vendor-proto:
		mkdir -p vendor.protogen
		mkdir -p vendor.protogen/api/ocp-presentation-api
		cp api/ocp-presentation-api/ocp-presentation-api.proto vendor.protogen/api/ocp-presentation-api
		@if [ ! -d vendor.protogen/google ]; then \
			git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
			mkdir -p  vendor.protogen/google/ &&\
			mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
			rm -rf vendor.protogen/googleapis ;\
		fi
		@if [ ! -d vendor.protogen/github.com/envoyproxy ]; then \
			mkdir -p vendor.protogen/github.com/envoyproxy &&\
			git clone https://github.com/envoyproxy/protoc-gen-validate vendor.protogen/github.com/envoyproxy/protoc-gen-validate ;\
		fi


.PHONY: deps
deps: install-go-deps

.PHONY: install-go-deps
install-go-deps: .install-go-deps

.PHONY: .install-go-deps
.install-go-deps:
		ls go.mod || go mod init
		go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
		go get -u github.com/golang/protobuf/proto
		go get -u github.com/golang/protobuf/protoc-gen-go
		go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
		go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
		tmpdir=$$(mktemp -d); cd $$tmpdir && export GO111MODULE=off \
		  && go get -d github.com/envoyproxy/protoc-gen-validate \
			&& cd $$GOPATH/src/github.com/envoyproxy/protoc-gen-validate && git checkout v0.1.0 \
			&& go build -o $$GOPATH/bin/protoc-gen-validate $$GOPATH/src/github.com/envoyproxy/protoc-gen-validate/main.go \
			&& cd -
