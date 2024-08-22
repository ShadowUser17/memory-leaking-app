FROM golang:1.22-alpine AS dev
WORKDIR /root
COPY ./ ./
RUN go mod tidy
RUN go build -ldflags="-s -w" -o ./app ./cmd/main.go

FROM alpine:latest
WORKDIR /root
COPY --from=dev /root/app /usr/local/bin/
ENTRYPOINT ["/usr/local/bin/app"]
