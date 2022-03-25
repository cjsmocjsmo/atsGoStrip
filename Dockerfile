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
  mkdir ./data && \
  mkdir ./data/db && \
  mkdir ./static && \
  mkdir ./static/assets && \
  mkdir ./static/assets/images && \
  mkdir ./backup && \
  mkdir ./fsData && \
  mkdir ./fsData/thumb && \
  mkdir ./fsData/crap && \
  mkdir ./logs

COPY backup/*.json ./backup/
COPY backup/*.gz ./backup/
COPY static/*.html ./static/
COPY static/*.yaml ./static/
COPY static/assets/*js ./static/assets/
COPY static/assets/*css ./static/assets/
COPY static/assets/images/*webp ./static/assets/images/
COPY static/assets/images/*jpg ./static/assets/images/

RUN \
  chmod -R +rwx ./static && \
  chmod -R +rwx ./fsData && \
  chmod -R +rwx ./logs && \
  chmod -R +rwx ./backup

STOPSIGNAL SIGINT
CMD ["./main"]

