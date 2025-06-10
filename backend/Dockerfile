FROM golang:1.24 as builder
WORKDIR /app
COPY . .
RUN go mod tidy && CGO_ENABLED=0 GOOS=linux go build -o vote ./cmd/server

FROM alpine
WORKDIR /
COPY --from=builder /app/vote .
EXPOSE 8080
CMD ["/vote"]
