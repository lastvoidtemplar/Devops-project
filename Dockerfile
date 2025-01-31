FROM golang:1.23-alpine AS builder

WORKDIR /build

COPY go.mod go.mod
COPY ./cmd .

RUN go build -o goapp .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /build/goapp .

EXPOSE 80

CMD [ "./goapp" ]