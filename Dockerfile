FROM golang:latest AS builder

COPY . /src
WORKDIR /src
RUN go mod download
RUN go build -o /out/ocp-presentation-api ./cmd/ocp-presentation-api

FROM alpine:latest

COPY --from=builder /out/ocp-presentation-api /ocp-presentation-api
EXPOSE 8080
CMD ["/ocp-presentation-api"]
