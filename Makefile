.PHONY: build
build: vendor-proto generate .build

.PHONY: .all
.all: dependencies build

.PHONY: all
all: requirements .all

PHONY: generate
generate:
		mkdir -p pkg/ocp-presentation-api
		mkdir -p pkg/ocp-slide-api
		mkdir -p swagger

		protoc -I vendor.protogen \
				--go_out=pkg/ocp-presentation-api --go_opt=paths=import \
				--go-grpc_out=pkg/ocp-presentation-api --go-grpc_opt=paths=import \
				--grpc-gateway_out=pkg/ocp-presentation-api \
				--grpc-gateway_opt=logtostderr=true \
				--grpc-gateway_opt=paths=import \
				--validate_out lang=go:pkg/ocp-presentation-api \
				--swagger_out=allow_merge=true,merge_file_name=ocp-presentation-api:swagger \
				api/ocp-presentation-api/ocp-presentation-api.proto
		mv pkg/ocp-presentation-api/gihtub.com/ozoncp/ocp-presentation-api/pkg/ocp-presentation-api/* pkg/ocp-presentation-api/
		rm -rf pkg/ocp-presentation-api/gihtub.com
		mkdir -p cmd/ocp-presentation-api

		protoc -I vendor.protogen \
				--go_out=pkg/ocp-slide-api --go_opt=paths=import \
				--go-grpc_out=pkg/ocp-slide-api --go-grpc_opt=paths=import \
				--grpc-gateway_out=pkg/ocp-slide-api \
				--grpc-gateway_opt=logtostderr=true \
				--grpc-gateway_opt=paths=import \
				--validate_out lang=go:pkg/ocp-slide-api \
				--swagger_out=allow_merge=true,merge_file_name=ocp-slide-api:swagger \
				api/ocp-slide-api/ocp-slide-api.proto
		mv pkg/ocp-slide-api/gihtub.com/ozoncp/ocp-presentation-api/pkg/ocp-slide-api/* pkg/ocp-slide-api/
		rm -rf pkg/ocp-slide-api/gihtub.com
		mkdir -p cmd/ocp-slide-api

		swagger mixin --ignore-conflicts swagger/ocp-presentation-api.swagger.json swagger/ocp-slide-api.swagger.json > swagger/api.swagger.json

PHONY: .build
.build:
ifeq ($(OS), Windows_NT)
	CGO_ENABLED=0 GOOS=windows go build -o bin/ocp-presentation-api.exe cmd/ocp-presentation-api/main.go
else
	CGO_ENABLED=0 GOOS=linux go build -o bin/ocp-presentation-api cmd/ocp-presentation-api/main.go
endif

.PHONY: requirements
requirements:
		apt-get update -q && apt-get install -y protobuf-compiler

.PHONY: dependencies
dependencies:
		ls go.mod || go mod init

		go get -u github.com/envoyproxy/protoc-gen-validate
		go get -u github.com/go-swagger/go-swagger/cmd/swagger
		go get -u github.com/golang/protobuf/proto
		go get -u github.com/golang/protobuf/protoc-gen-go
		go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
		go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
		go get -u github.com/rs/zerolog/log
		go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
		go get -u google.golang.org/protobuf/runtime/protoimpl
		go get -u google.golang.org/protobuf/types/known/anypb

		go mod download

		go install github.com/envoyproxy/protoc-gen-validate
		go install github.com/go-swagger/go-swagger/cmd/swagger
		go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
		go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

.PHONY: vendor-proto
vendor-proto:
		mkdir -p vendor.protogen/api/ocp-presentation-api
		mkdir -p vendor.protogen/api/ocp-slide-api

		cp -rf api/ocp-presentation-api/ocp-presentation-api.proto vendor.protogen/api/ocp-presentation-api/
		cp -rf api/ocp-slide-api/ocp-slide-api.proto vendor.protogen/api/ocp-slide-api/

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

.PHONY: coverage
coverage:
	go test -race -coverprofile="coverage.out" -covermode=atomic ./...
	go tool cover -html="coverage.out"

.PHONY: clean
clean:
		rm -rf bin pkg swagger vendor.protogen coverage.out

.PHONY: tidy
tidy:
		go mod tidy

.PHONY: lint
lint:
		golangci-lint run

.PHONY: grpcui
grpcui:
		grpcui -proto ./api/ocp-presentation-api/ocp-presentation-api.proto -import-path ./vendor.protogen -plaintext -open-browser localhost:8000

.PHONY: deploy
deploy:
		docker-compose up -d
