FROM golang:alpine as build-env

ENV GO111MODULE=on

RUN apk update && apk add bash ca-certificates git gcc

RUN mkdir /chat-docker-grpc
RUN mkdir -p /chat-docker-grpc/proto

WORKDIR /chat-docker-grpc

COPY ./proto/service.pb.go /chat-docker-grpc/proto
COPY ./main.go /chat-docker-grpc

COPY go.mod .
COPY go.sum .

RUN go mod download

RUN go build -o chat-docker-grpc .

CMD ./chat-docker-grpc