FROM golang:1.22.3-alpine3.16 as builder
WORKDIR /app
RUN apk update && \
    apk add --no-cache \
    coreutils \
    git \
    make
COPY . .
RUN make build

FROM alpine:3.11
COPY --from=builder /app/bin/linux_amd64/shortener /usr/bin
CMD ["/usr/bin/shortener"]