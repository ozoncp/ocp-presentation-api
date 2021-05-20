FROM golang:latest AS builder

COPY . /src
WORKDIR /src
RUN go get -d -v ./...
RUN go build -o /out/app ./cmd/ocp-presentation-api

FROM alpine:latest

COPY --from=builder /out/app /app
CMD ["/app"]
