FROM golang:1.17.2-buster as builder

WORKDIR /app

RUN apt-get install -y ca-certificates

COPY ["go.mod", "go.sum", "./"]
RUN go mod download

COPY cmd cmd
COPY internal internal

COPY Makefile .

RUN make build

FROM debian:buster

WORKDIR /app
USER nobody

COPY --from=builder /app/server /app/server
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 8000

ENTRYPOINT ["/app/server"]

