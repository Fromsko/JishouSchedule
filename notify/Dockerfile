FROM alpine:latest

# 设置工作目录
RUN mkdir -p /code
WORKDIR /code

# 复制程序
Copy ./autoTask/notify /code/notify

# 重设时区
RUN apk --update add tzdata && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone && \
    apk del tzdata && \
    rm -rf /var/cache/apk/*

# 设置运行权限
Run chmod +x ./notify

# 设置容器启动时的默认命令
CMD ["./notify"]