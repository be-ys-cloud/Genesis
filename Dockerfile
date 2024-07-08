FROM golang:1.22-alpine AS builder

WORKDIR /go/src/genesis
RUN mkdir -p /app

COPY . .
RUN go mod download
RUN go build -o /app/genesis ./cmd

FROM alpine:3.20
COPY --from=builder /app/genesis /app/genesis

USER nobody
EXPOSE 8080
HEALTHCHECK --interval=30s --timeout=5s --start-period=10s \
	CMD wget -q --spider "http://localhost:8080/metrics"

WORKDIR /app
ENTRYPOINT ["./genesis"]
