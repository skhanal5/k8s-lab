FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o k8s-lab ./cmd/k8s-lab


FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/k8s-lab .

EXPOSE 8080

CMD ["./k8s-lab"]