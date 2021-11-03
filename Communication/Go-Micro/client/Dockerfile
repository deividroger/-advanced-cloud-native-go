FROM golang:1.13.0-alpine

RUN apk update && apk upgrade && apk add --no-cache bash git gcc

RUN export GO111MODULE=on

ENV SOURCES /go/src/github.com/deividroger/advanced-cloud-native-go/Communication/Go-Micro/
COPY . ${SOURCES}

RUN cd ${SOURCES}/client  && \
     go get google.golang.org/protobuf/cmd/protoc-gen-go \
        google.golang.org/grpc/cmd/protoc-gen-go-grpc && \

     go install && \
     CGO_ENABLED=0 go build

ENV CONSUL_HTTP_ADDR localhost:8500

WORKDIR ${SOURCES}client/
CMD ${SOURCES}client/client