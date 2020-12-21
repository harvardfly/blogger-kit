FROM golang:1.14:latest
WORKDIR /src
COPY / /home/zpf/go/src/blogger-kit
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go build -o user
EXPOSE 10086
ENTRYPOINT ./user -mysql.addr $mysqlAddr -redis.addr $redisAddr