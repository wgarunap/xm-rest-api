FROM golang:1.16-alpine3.14 AS builder

WORKDIR /src
COPY . .
RUN go build -ldflags="-s -w" -o ./xm-rest-api *.go

FROM alpine:3.14.0

WORKDIR /opt
COPY --from=builder /src/xm-rest-api /opt

ENTRYPOINT ["sh","-c","./xm-rest-api"]



