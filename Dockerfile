
FROM golang:1.23 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o simswap main.go

FROM debian:bookworm-slim

WORKDIR /root/

COPY --from=builder /app/simswap .

EXPOSE 9091

CMD ["./simswap"]
