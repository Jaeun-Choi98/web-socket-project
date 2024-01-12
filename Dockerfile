# Stage 1: Build Go 백엔드
FROM golang:latest as go-build
WORKDIR /chat-server
COPY ./ ./
RUN go mod download && CGO_ENABLED=0 GOOS=linux go build -o chatserver

# Stage 2: 경량 리눅스로 실행
FROM alpine:latest
WORKDIR /app
COPY --from=go-build /chat-server/chatserver /app/
CMD ["./chatserver"]
