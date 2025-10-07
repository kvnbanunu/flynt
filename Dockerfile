FROM golang:1.25.1-alpine AS builder

WORKDIR /app

RUN apk add --no-cache gcc musl-dev sqlite-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -o main .

FROM alpine:latest

RUN apk --no-cache add ca-certificates sqlite-libs

WORKDIR /root/

RUN mkdir -p /data

COPY --from=builder /app/main .

COPY .env .env

EXPOSE 8000

CMD ["./main"]
