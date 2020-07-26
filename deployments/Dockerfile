FROM golang:alpine as build

ENV GOPROXY=https://goproxy.io

ADD . /gingo

WORKDIR /gingo

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o gingo_server cmd/main.go

FROM alpine:3.7

RUN echo "http://mirrors.aliyun.com/alpine/v3.7/main/" > /etc/apk/repositories && \
    apk update && \
    apk add ca-certificates && \
    mkdir -p /www/configs && \
    mkdir -p /www/logs

WORKDIR /www
ENV PATH="/www:${PATH}"

COPY --from=build /gingo/gingo_server /www/gingo_server
COPY --from=build /gingo/configs/conf.yaml /www/configs/conf.yaml

RUN chmod +x /www/gingo_server

ENTRYPOINT ["gingo_server"]
