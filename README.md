# Ozon Code Platform Presentation API
[![Build & Test](https://github.com/ozoncp/ocp-presentation-api/actions/workflows/workflow.yml/badge.svg?branch=main)](https://github.com/ozoncp/ocp-presentation-api/actions/workflows/workflow.yml)
[![codecov](https://codecov.io/gh/ozoncp/ocp-presentation-api/branch/main/graph/badge.svg?token=sjlJtE7Yb1)](https://codecov.io/gh/ozoncp/ocp-presentation-api)

Ozon Code Platform Presentation API is a service that provides API to access/manage Presentation and Slide entities.

## Go Generated Code

```sh
make requirements
make dependencies
make vendor-proto
make generate
```

## [swagger-ui](http://localhost:80/swagger)

### Download
```sh
docker pull swaggerapi/swagger-ui
```
### Run
On Unix
```sh
docker run -d -p 80:8080 -e BASE_URL=/swagger -e SWAGGER_JSON=/swagger/api.swagger.json -v `pwd`/swagger:/swagger swaggerapi/swagger-ui
```
On Windows

```cmd
docker run -d -p 80:8080 -e BASE_URL=/swagger -e SWAGGER_JSON=/swagger/api.swagger.json -v %cd%/swagger:/swagger swaggerapi/swagger-ui
```

## grpcui
```sh
grpcui -proto ./api/ocp-presentation-api/ocp-presentation-api.proto -import-path ./vendor.protogen -plaintext -open-browser localhost:8000
```

# docker-compose
```sh
docker-compose pull

docker-compose run ocp-presentation-api

docker-compose up -d
docker-compose build
docker-compose up --build

docker-compose ps
docker logs --follow ocp-presentation-api_web_1

docker-compose stop
docker-compose down --volumes
```

## [prometheus](http://localhost:9090)

## [Jaeger UI](http://localhost:16686)
