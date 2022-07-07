FROM golang:1.18-alpine as builder

WORKDIR /root

COPY . ./
RUN go mod download && \
    go mod verify && \
    go build -o ./build/app ./cmd/app/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /root/build/app ./app

EXPOSE 8080 8081

CMD ["./app"]