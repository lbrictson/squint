# Build in app in a Go container
FROM docker.io/golang:1.17.3-buster as builder
COPY .  /app
WORKDIR /app
RUN go env -w GOPROXY=direct && go env -w GOSUMDB=off && CGO_ENABLED=0 go build -o main cmd/main.go
# Move artifact to smaller container with no Go tools installed
FROM docker.io/alpine:3.15.0
RUN apk update && apk upgrade && apk --no-cache add ca-certificates bash
WORKDIR /app
COPY --from=builder /app/main app
ENTRYPOINT ["/app/app"]