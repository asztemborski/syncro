FROM golang:1.23-alpine as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

ARG TARGETARCH
RUN GOARCH=$TARGETARCH go build -o main .

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=builder /app/config /root/config/

EXPOSE 8080

CMD ["./main"]