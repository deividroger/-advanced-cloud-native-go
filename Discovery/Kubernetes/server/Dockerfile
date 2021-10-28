FROM golang:1.8.1-alpine

ENV SOURCES /go/src/github.com/deividroger/Advanced-Cloud-Native-Go/Discovery/Kubernetes/
COPY . ${SOURCES}

RUN cd ${SOURCES}server/ && CGO_ENABLED=0 go build

WORKDIR ${SOURCES}server/
CMD ${SOURCES}server/server