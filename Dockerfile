FROM golang:1.17 AS builder

ADD . /app
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
RUN go build -o /app

FROM alpine:latest AS production
COPY --from=builder /app .
CMD ["./app"]