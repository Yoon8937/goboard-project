FROM golang:1.26-alpine AS builder
RUN apk update && apk add --no-cache git tzdata
WORKDIR /app
COPY go.mod go.sum ./


RUN go mod download
#RUN --mount=type=cache,target=/go/pkg/mod \
#    --mount=type=bind,source=go.mod,target=go.mod \
#    --mount=type=bind,source=go.sum,target=go.sum \
#    go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o server ./cmd/server/main.go
#RUN --mount=type=cache,target=/go/pkg/mod \
#    --mount=type=cache,target=/root/.cache/go-build \
#    CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o server ./cmd/server/main.go



FROM alpine:3.19
RUN apk update && apk add --no-cache tzdata ca-certificates
ENV TZ=Asia/Seoul
WORKDIR /app
COPY --from=builder /app/server .
COPY --from=builder /app/templates ./templates
EXPOSE 8080
CMD ["./server"]
