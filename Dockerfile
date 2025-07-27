FROM golang:1.22-alpine AS builder
LABEL maintainer="awlhajb@gmail.com"
LABEL project="ascii-art-web"

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .
RUN go build -o server

FROM alpine:latest

LABEL stage="production"
LABEL version="1.0"

WORKDIR /app
COPY --from=builder /app/server .
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/static ./static


EXPOSE 8080

CMD ["./server"]