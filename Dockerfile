# 第一阶段：构建 Go 二进制文件
FROM golang:1.23-alpine AS builder

# 设置环境变量，确保 Go 模块被启用
ENV GO111MODULE=on

# 设置工作目录为 /app，而不是 $GOPATH/src
WORKDIR /app

# 复制 go.mod 和 go.sum 到工作目录
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制所有源代码到工作目录
COPY . .

# 编译 Go 应用，生成二进制文件
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o url_shortener

# 第二阶段：创建最小化运行环境
FROM alpine:latest

# 安装必要的运行时库（例如证书）
RUN apk --no-cache add ca-certificates

# 设置工作目录
WORKDIR /root/

# 从构建阶段复制二进制文件
COPY --from=builder /app/url_shortener .

# 暴露应用运行的端口
EXPOSE 8080

# 健康检查（可选）
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
  CMD wget -qO- http://localhost:8080/ || exit 1

# 启动应用
CMD ["./url_shortener"]