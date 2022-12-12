FROM golang:alpine AS builder
ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct
WORKDIR /build
COPY . .
RUN go mod download
RUN go build -ldflags="-s -w" -o easy-cron ./main.go


FROM harbor-tj.ts-it.cn:63333/bigdata/flinkx-dist:1.12.5
WORKDIR /app
ENV TZ Asia/Shanghai
MAINTAINER Rethink
#更新Alpine的软件源为国内（清华大学）的站点，因为从默认官源拉取实在太慢了。。。
RUN echo "https://mirror.tuna.tsinghua.edu.cn/alpine/v3.4/main/" > /etc/apk/repositories
RUN apk update \
        && apk upgrade \
        && apk add --no-cache bash \
        bash-doc \
        bash-completion \
        && rm -rf /var/cache/apk/* \
        && /bin/bash
COPY --from=builder /build/config.yaml /app/config.yaml
COPY --from=builder /build/easy-cron /app/easy-cron
CMD ["./easy-cron"]