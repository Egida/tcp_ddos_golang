# Compile app binary
FROM golang:1.20.1-alpine3.16 as build-env

WORKDIR /go/src/github.com/ypapax/tcp_ddos_golang
ENV GO111MODULE=on
RUN apk add git
COPY go.mod go.mod
COPY go.sum go.sum

COPY hashcash2 hashcash2
COPY apps/client apps/client

WORKDIR apps/client
RUN ls -laht
RUN pwd

RUN go get ./...
RUN GOOS=linux GOARCH=amd64 go install

CMD ["client"]
#ENTRYPOINT ["bash", "/entrypoint.sh"]