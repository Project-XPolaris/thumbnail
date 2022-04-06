ARG GOLANG_VERSION=1.17
FROM golang:${GOLANG_VERSION}-alpine as builder
ARG CGO_CFLAGS_ALLOW="-Xpreprocessor"
ARG GOPROXY=https://goproxy.cn
RUN apk add pkgconfig git gcc musl-dev vips-dev

WORKDIR ${GOPATH}/src/github.com/projectxpolaris/thumbnailservice

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ${GOPATH}/bin/thumbnailservice ./main.go

FROM alpine
ARG GOPROXY=https://goproxy.cn
COPY --from=builder /usr/local/lib /usr/local/lib
COPY --from=builder /etc/ssl/certs /etc/ssl/certs

RUN apk add vips-dev

COPY --from=builder /go/bin/thumbnailservice /usr/local/bin/thumbnailservice

ENV VIPS_WARNING=0
ENV MALLOC_ARENA_MAX=2


# use unprivileged user
USER nobody

ENTRYPOINT ["/usr/local/bin/thumbnailservice","run"]