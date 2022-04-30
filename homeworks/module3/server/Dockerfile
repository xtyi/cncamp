FROM golang:1.18.1 AS builder

COPY . /src
WORKDIR /src

RUN go build .

FROM alpine:3.15.4

COPY --from=builder /src/server /app

WORKDIR /app

EXPOSE 8080

CMD [ "./server" ]
