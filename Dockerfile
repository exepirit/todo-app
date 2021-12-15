# [Stage 1] - Build server
FROM golang:1.16-alpine AS builder

WORKDIR /build
COPY . .

RUN go mod download
# RUN go test ./...
RUN go build ./cmd/todoserver

# [Stage 2] - Artifact image
FROM alpine

COPY --from=builder /build/todoserver /usr/bin/todoserver
CMD ["/usr/bin/todoserver"]
