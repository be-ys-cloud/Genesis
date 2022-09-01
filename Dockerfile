FROM golang:1.19-alpine

WORKDIR /go/src/genesis
RUN mkdir -p /app


COPY . .
RUN go mod download
RUN go build -o /app/genesis ./cmd

FROM golang:alpine
COPY --from=0 /app/genesis /app/genesis

EXPOSE 8080

WORKDIR /app
ENTRYPOINT ["./genesis"]

