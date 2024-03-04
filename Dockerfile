FROM golang:1.20.3-alpine AS builder

COPY . /github.com/kenyako/auth/source/
WORKDIR /github.com/kenyako/auth/source/

RUN go mod download
RUN go mod tidy -e
RUN go build -o ./bin/server cmd/grpc_server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/kenyako/auth/source/bin/server .

CMD [ "./server" ]