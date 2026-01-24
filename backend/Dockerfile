# ===== 构建阶段 =====
FROM golang:1.25-alpine AS builder

WORKDIR /app

ENV GOPROXY=https://goproxy.cn,direct
ENV GO111MODULE=on

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o blog ./cmd


# ===== 运行阶段 =====
FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/blog /app/blog
COPY config /app/config
COPY uploads /app/uploads

EXPOSE 8080

ENV GIN_MODE=release

CMD ["/app/blog"]
