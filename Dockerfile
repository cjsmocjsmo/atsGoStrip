FROM golang:bullseye AS builder
RUN mkdir /go/src/atsGo
WORKDIR /go/src/atsGo

COPY atsGo.go .

COPY go.mod .
COPY go.sum .
RUN export GOPATH=/go/src/atsGo
RUN go get -v /go/src/atsGo
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main /go/src/atsGo

# FROM ubuntu:22.04
# FROM alpine:latest

# ARG DEBIAN_FRONTEND=noninteractive
# ENV TZ=America/New_York

# RUN \
#     apt-get update && \
#     apt-get -y dist-upgrade && \
#     apt-get -y install golang openssl libssl-dev && \
#     apt-get -y autoclean && \
#     apt-get -y autoremove  

WORKDIR /root/

COPY --from=builder /go/src/atsGo/main .

RUN \
  mkdir ./data && \
  mkdir ./data/db && \
  mkdir ./assets && \
  mkdir ./backup 

COPY backup/*.json ./backup/
COPY backup/*.gz ./backup/
COPY assets/*.html ./assets/
COPY assets/*.yaml ./assets/

RUN \
  chmod -R +rwx ./assets && \
  chmod -R +rwx ./backup

STOPSIGNAL SIGINT
CMD ["./main"]

