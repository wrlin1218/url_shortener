# 第一阶段：构建 Go 二进制文件
FROM golang:1.20-alpine AS builder

WORKDIR .

# 仅复制 go.mod 和 go.sum 文件以利用缓存
COPY go.mod go.sum ./
RUN go mod download

# 复制所有源代码
COPY . .

# 编译 Go 应用
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o url_shortener

# 第二阶段：创建最小化运行环境
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder ./url_shortener .

EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
  CMD wget -qO- http://localhost:8080/ || exit 1

CMD ["./url_shortener"]