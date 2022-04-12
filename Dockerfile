# FROM node:14bullseye-slim AS nodebuilder
# WORKDIR /home/node
# RUN \
#   apt-get -y install git nodejs npm && \
#   cd /root && \
#   git clone http://github.com/cjsmocjsmo/AlphaTreeService && \
#   cd AlphaTreeService && \
#   npm install && \
#   npm run build



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
# FROM cjsmocjsmo/alpha:0.0.1

WORKDIR /root/

COPY --from=builder /go/src/atsGo/main .

RUN \
  mkdir ./data && \
  mkdir ./data/db && \
  mkdir ./assets && \
  # mkdir ./assets/images && \
  mkdir ./backup && \
  mkdir ./uploads && \
  # mkdir ./fsData && \
  # mkdir ./fsData/thumb && \
  # mkdir ./fsData/crap && \
  mkdir ./logs

COPY backup/*.json ./backup/
COPY backup/*.gz ./backup/
COPY assets/*.html ./assets/
COPY assets/*.yaml ./assets/
COPY assets/*js ./assets/


# COPY assets/*css ./assets/
# COPY assets/images/*webp ./assets/images/
# COPY assets/images/*jpg ./assets/images/

RUN \
  chmod -R +rwx ./assets && \
  # chmod -R +rwx ./fsData && \
  chmod -R +rwx ./logs && \
  chmod -R +rwx ./backup

STOPSIGNAL SIGINT
CMD ["./main"]

