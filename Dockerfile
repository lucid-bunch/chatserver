FROM golang:1.12-alpine as builder
WORKDIR /build/
ADD . .
RUN apk add --update git && \
    go mod download && \
    CGO_ENABLED=0 GOOS=linux go build -ldflags '-s -w -extldflags "-static"' -o runnable_app

FROM scratch
USER 1000
COPY --from=builder /build/runnable_app .
CMD ["./app"]
