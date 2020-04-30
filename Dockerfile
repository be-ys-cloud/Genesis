FROM golang:alpine

RUN apk --update add git

RUN git clone https://github.com/be-ys/Genesis.git /go/src/genesis

WORKDIR /go/src/genesis

ENV GO111MODULE=on
RUN go mod init
WORKDIR /go/src/genesis/src

RUN mkdir -p /app
RUN go build -o /app/genesis .
RUN rm -rf /go/src/genesis

COPY config.json /app/config.json

EXPOSE 8080

WORKDIR /app
ENTRYPOINT ["./genesis"]

