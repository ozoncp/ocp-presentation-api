FROM golang:latest AS builder

COPY . /build
WORKDIR /build
RUN make all

FROM alpine:latest AS ocp-presentation-api

COPY --from=builder /build/bin/ocp-presentation-api /ocp-presentation-api
CMD ["/ocp-presentation-api"]

FROM alpine:latest AS ocp-slide-api

COPY --from=builder /build/bin/ocp-slide-api /ocp-slide-api
CMD ["/ocp-slide-api"]
