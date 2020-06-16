FROM golang:alpine as build

ENV GOPROXY=https://goproxy.io

ADD . /gingo

WORKDIR /gingo

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o gingo_server

FROM alpine:3.7

RUN echo "http://mirrors.aliyun.com/alpine/v3.7/main/" > /etc/apk/repositories && \
    apk update && \
    apk add ca-certificates && \
    echo "hosts: files dns" > /etc/nsswitch.conf && \
    mkdir -p /www/conf && \
    mkdir -p /www/logs

WORKDIR /www
ENV PATH="/www:${PATH}"

COPY --from=build /gingo/gingo_server /www/gingo_server
COPY --from=build conf.yaml /www/conf.yaml

RUN chmod +x /www/gingo_server

ENTRYPOINT ["gingo_server"]
