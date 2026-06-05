FROM golang:1.26-bookworm AS builder

RUN apt-get update && apt-get install -y gcc g++

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -o pitsinkr ./cmd/server

FROM debian

RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY --from=builder /app/pitsinkr .

EXPOSE 53/udp
EXPOSE 4004
EXPOSE 9056

CMD ["./pitsinkr"]