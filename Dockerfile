FROM golang:1.18-alpine3.14 AS build
COPY . /opt
WORKDIR /opt

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories \
    && apk update && apk add git gcc \
    && GOPROXY=https://goproxy.cn,direct go build -o Furbot-MiraiGo

FROM alpine:3.14
COPY --from=build /opt/Furbot-MiraiGo /opt/Furbot-MiraiGo
WORKDIR /opt

CMD ["./Furbot-MiraiGo"]
