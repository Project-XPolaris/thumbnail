ARG GOLANG_VERSION=1.17
FROM golang:${GOLANG_VERSION}-buster as builder

ARG GOPROXY=https://goproxy.cn
# Installs libvips + required libraries
RUN DEBIAN_FRONTEND=noninteractive \
  apt-get update && \
  apt-get install --no-install-recommends -y \
  libvips-dev

WORKDIR ${GOPATH}/src/github.com/projectxpolaris/thumbnailservice

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ${GOPATH}/bin/thumbnailservice ./main.go

FROM debian:buster-slim
ARG GOPROXY=https://goproxy.cn
COPY --from=builder /usr/local/lib /usr/local/lib
COPY --from=builder /etc/ssl/certs /etc/ssl/certs

# Install runtime dependencies
RUN DEBIAN_FRONTEND=noninteractive \
  apt-get update && \
  apt-get install --no-install-recommends -y \
  libvips-dev

COPY --from=builder /go/bin/thumbnailservice /usr/local/bin/thumbnailservice

ENV VIPS_WARNING=0
ENV MALLOC_ARENA_MAX=2


# use unprivileged user
USER nobody

ENTRYPOINT ["/usr/local/bin/thumbnailservice","run"]