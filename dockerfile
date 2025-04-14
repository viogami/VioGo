FROM alpine:latest

# 把本地的可执行文件复制进镜像
COPY main /app/main
COPY config.yaml /app/config.yaml

# 设置工作目录
WORKDIR /app

# 启动命令
CMD ["./main"]
