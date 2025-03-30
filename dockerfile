FROM golang:1.24 AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server ./main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/server .

EXPOSE 1011

CMD ["./server"]