FROM golang:alpine AS builder

WORKDIR /app
COPY ./ ./

WORKDIR /app/cmd/tempestas
RUN go build -o tempestas

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/cmd/tempestas/tempestas .

EXPOSE 8000
ENTRYPOINT ["./tempestas"]

