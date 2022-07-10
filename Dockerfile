FROM golang:1.12.0-alpine3.9 AS builder

ENV GO111MODULES=on

WORKDIR /go_worlder_system
COPY ./ /go_worlder_system

RUN apk add --no-cache ca-certificates git && \
    go build -o worlder_system


FROM alpine:3.9

RUN apk add --no-cache ca-certificates

WORKDIR /go_worlder_system
COPY --from=builder /go_worlder_system /go_worlder_system

EXPOSE 8080
ENTRYPOINT ["/go_worlder_system/worlder_system"]
