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

## [swagger-ui](https://editor.swagger.io/)
## [Local swagger-ui](http://localhost:80/swagger)

### Download
```sh
docker pull swaggerapi/swagger-ui
```
### Run
on Unix
```sh
docker run -d -p 80:8080 -e BASE_URL=/swagger -e SWAGGER_JSON=/swagger/api.swagger.json -v `pwd`/swagger:/swagger swaggerapi/swagger-ui
```
On Windows

```cmd
docker run -d -p 80:8080 -e BASE_URL=/swagger -e SWAGGER_JSON=/swagger/api.swagger.json -v %cd%/swagger:/swagger swaggerapi/swagger-ui
```

## grpcui
```sh
grpcui -proto ./api/ocp-presentation-api/ocp-presentation-api.proto -import-path ./vendor.protogen -plaintext -open-browser localhost:7002
```
