FROM golang:1.13.0-alpine

RUN apk update && apk upgrade && apk add --no-cache bash git

ENV SOURCES /go/src/github.com/deividroger/Advanced-Cloud-Native-Go/Discovery/Sample/
COPY . ${SOURCES}

RUN cd ${SOURCES}/client &&  \
    go install && \
    CGO_ENABLED=0 go build


#RUN cd ${SOURCES}client/ && CGO_ENABLED=0 go build

ENV CONSUL_HTTP_ADDR localhost:8500

WORKDIR ${SOURCES}client/
CMD ${SOURCES}client/client