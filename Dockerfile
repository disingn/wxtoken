# 使用官方Go镜像作为构建环境
FROM golang:1.20 as builder

# 设置工作目录
WORKDIR /app

# 复制go.mod和go.sum文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 构建应用程序并设置执行权限
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o wxToken .
#RUN #chmod +x /app/wxToken

FROM alpine

RUN apk update \
    && apk upgrade \
    && apk add --no-cache ca-certificates tzdata \
    && update-ca-certificates 2>/dev/null || true

# 从构建器镜像中复制执行文件
COPY --from=builder /app/wxToken /wxToken

# 暴露端口
EXPOSE 3100

# 运行应用程序
CMD ["/wxToken"]
