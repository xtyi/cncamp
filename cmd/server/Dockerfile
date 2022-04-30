FROM golang:1.18.1 AS builder

COPY . /src
WORKDIR /src

RUN CGO_ENABLED=0 go build -o server main.go

FROM alpine:3.15.4

COPY --from=builder /src/server /app/server

WORKDIR /app

EXPOSE 8080

CMD [ "./server" ]
