ARG IMAGE=scratch
ARG OS=linux
ARG ARCH=amd64

FROM golang:1.23.6-alpine AS builder

WORKDIR /go/src/github.com/soerenuhrbach/egym-exporter
COPY . .

RUN apk --no-cache add git alpine-sdk

RUN GO111MODULE=on go mod vendor
RUN CGO_ENABLED=0 GOOS=$OS GOARCH=$ARCH go build -ldflags '-s -w' -o binary ./

FROM $IMAGE

LABEL name="egym-exporter"

WORKDIR /root/
COPY --from=builder /go/src/github.com/soerenuhrbach/egym-exporter/binary egym-exporter

CMD ["./egym-exporter"]