FROM golang:1.23-alpine as builder

WORKDIR /app

RUN export GO111MODULE=on

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 go build -v -o ./walle ./cmd

FROM alpine:latest

COPY --from=builder /app/walle .
COPY --from=builder /app/config/config.json .

ENTRYPOINT ["./walle"]