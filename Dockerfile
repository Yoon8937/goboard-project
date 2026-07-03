FROM golang:1.26-alpine AS builder
RUN apk update && apk add --no-cache git tzdata
WORKDIR /app
COPY go.mod go.sum ./


RUN go mod download
COPY . .

# Go 애플리케이션 빌드
# CGO_ENABLED=0 : 순수 Go 바이너리로 빌드하여 다른 리눅스 환경에서도 독립 실행 가능하게 함
# GOOS=linux : 타겟 OS를 리눅스로 지정
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o server ./cmd/server/main.go




FROM alpine:3.19
RUN apk update && apk add --no-cache tzdata ca-certificates
ENV TZ=Asia/Seoul
WORKDIR /app
COPY --from=builder /app/server .
COPY --from=builder /app/templates ./templates
EXPOSE 8080
CMD ["./server"]
