FROM golang:latest AS builder

COPY . /ocp-presentation-api
WORKDIR /ocp-presentation-api
RUN make requirements && make dependencies && make build

FROM alpine:latest

COPY --from=builder /ocp-presentation-api/bin/ocp-presentation-api /ocp-presentation-api
CMD ["/ocp-presentation-api"]
