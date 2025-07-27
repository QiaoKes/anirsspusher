# 构建阶段
FROM golang:1.23-alpine AS builder

# 设置工作目录
WORKDIR /app

# 复制依赖文件并下载依赖
COPY go.mod go.sum ./
RUN go mod tidy

# 复制所有源代码
COPY . .

RUN  go build -o app .

# ----------------------------------------
# 运行时阶段 (使用极简基础镜像)
FROM alpine:3.18

# 安装时区数据（可选）
RUN apk add --no-cache tzdata

# 从构建阶段复制二进制文件
COPY --from=builder /app/app /app
COPY --from=builder /app/conf /conf

# 暴露端口
EXPOSE 8080

# 设置配置文件挂载点
VOLUME ["/conf"]

# 设置容器启动命令
CMD ["/app"]