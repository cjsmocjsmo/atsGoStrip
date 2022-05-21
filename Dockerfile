FROM golang:bullseye AS builder
RUN mkdir /go/src/atsGo
WORKDIR /go/src/atsGo

COPY atsGo.go .

COPY go.mod .
COPY go.sum .
RUN export GOPATH=/go/src/atsGo
RUN go get -v /go/src/atsGo
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main /go/src/atsGo

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /go/src/atsGo/main .

RUN \
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

