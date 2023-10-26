# syntax=registry.cn-hangzhou.aliyuncs.com/lcc-middleware/docker-dockerfile:1.4
FROM registry.cn-hangzhou.aliyuncs.com/0x0034/golang:1.20-alpine3.17 AS builder
# 配置alpine使用清华源
RUN echo -e "http://mirrors.aliyun.com/alpine/v3.17/main\nhttp://mirrors.aliyun.com/alpine/v3.17/community" > /etc/apk/repositories && \
    apk update && apk add --no-cache git gcc musl-dev
WORKDIR /home/youngdev-stack/minio-ffmpeg-lambda/
# 将当前目录（也就是执行docker build xxx 的目录）下的文件拷贝到工作目录的上级目录下
# 配置docker镜像的系统环境：启用go modules
RUN go env -w GO111MODULE=on && \
    go env -w GOPROXY=https://goproxy.cn,direct
COPY  ./go.mod .
COPY  ./go.sum .
RUN go mod download -x
COPY . .
# 构建项目，输出到镜像的指定目录下
RUN go build -ldflags="-w -s" -o minio-ffmpeg-lambda ./

# 使用一个空的镜像打包发布后的go项目，以达到镜像体积的最小化
FROM registry.cn-hangzhou.aliyuncs.com/0x0034/ffmpeg:6.0-alpine

ENV LANG en_US.UTF-8
# 从上面的镜像中拷贝编译后的程序到当前镜像x的指定位置
WORKDIR /home/youngdev-stack/minio-ffmpeg-lambda/
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone
RUN echo -e "http://mirrors.aliyun.com/alpine/v3.13/main\nhttp://mirrors.aliyun.com/alpine/v3.13/community" > /etc/apk/repositories && \
    apk --no-cache add tzdata bash coreutils && \
    chmod u+s /bin/ping && \
    rm -rf /var/cache/apk/*

COPY --from=builder --chmod=755 /home/youngdev-stack/minio-ffmpeg-lambda/minio-ffmpeg-lambda .
COPY --from=builder --chmod=755 /home/youngdev-stack/minio-ffmpeg-lambda/docker-entrypoint.sh /usr/local/bin/
COPY --from=builder /home/youngdev-stack/minio-ffmpeg-lambda/config.yaml /home/youngdev-stack/minio-ffmpeg-lambda/conf/config.yaml

ENTRYPOINT ["docker-entrypoint.sh"]
EXPOSE 8888
