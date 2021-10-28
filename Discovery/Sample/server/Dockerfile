FROM golang:1.13.0-alpine

RUN apk update && apk upgrade && apk add --no-cache bash git

#RUN go get "github.com/hashicorp/consul/api"

ENV SOURCES /go/src/github.com/deividroger/Advanced-Cloud-Native-Go/Discovery/Sample/
COPY . ${SOURCES}
#RUN go get ${SOURCES}/.

RUN cd ${SOURCES}/server &&  \
    go install && \
    CGO_ENABLED=0 go build

#RUN cd ${SOURCES}server/ && CGO_ENABLED=0 go build

ENV CONSUL_HTTP_ADDR localhost:8500

WORKDIR ${SOURCES}server/
CMD ${SOURCES}server/server