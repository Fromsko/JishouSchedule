# 微信测试号通知

> 搭配 [JishouSchedule](https://github.com/Fromsko/JishouSchedule) 使用

## 🚀 快速开始

1. 拉取项目

   ```shell
   git clone git@github.com:Fromsko/JishouSchedule.git
   cd JishouSchedule/notify
   go mod tidy
   go install github.com/rakyll/statik
   ```

2. 构建程序

   ```shell
   # 更新依赖文件
   ./bin/update

   # 构建
   make

   # 清理
   make clean
   ```

3. 填写配置文件 `JishouSchedule/notify/config.yaml`

   ```shell
   vim config.yaml
   ```

### 正常启动

```shell
chmod +x ./notify
./notify
```

### Docker 启动

> 项目配置了 Docker 可一键部署执行

#### 构建文件

```Dockerfile
FROM alpine:latest

# 设置工作目录
RUN mkdir -p /code
WORKDIR /code

# 复制程序
Copy ./notify/build/linux_amd64/linux_amd64 /code/notify

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
```

#### 启动命令

```shell
docker run -itd --restart=always -v /code/JishouSchedule/notify/:/code/ --name=notify_sche notify:2.4
```

## ©️ 许可

本项目基于 MIT 许可证，请查阅 LICENSE 文件以获取更多信息。
