FROM golang:1.20.1-alpine3.16

WORKDIR /go/src/github.com/ypapax/tcp_ddos_golang
ENV GO111MODULE=on
RUN apk add git
COPY go.mod go.mod
COPY go.sum go.sum

COPY apps/server apps/server

WORKDIR apps/server
RUN ls -laht
RUN pwd

RUN go get ./...
RUN GOOS=linux GOARCH=amd64 go install

EXPOSE 9001
CMD ["server"]