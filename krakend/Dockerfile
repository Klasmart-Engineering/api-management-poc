FROM golang:1.17-alpine AS builder

RUN apk add make gcc musl-dev

WORKDIR /tmp/builder
RUN mkdir plugin
COPY plugin plugin
WORKDIR /tmp/builder/plugin

RUN go build -buildmode=plugin -o krakend-debugger.so .

FROM devopsfaith/krakend as krakend

COPY krakend.json /etc/krakend/krakend.json

USER root
RUN mkdir -p /etc/krakend/plugins
COPY  --from=builder /tmp/builder/plugin/krakend-debugger.so /etc/krakend/plugins/krakend-debugger.so
RUN chmod +x /etc/krakend/plugins/*.so

USER krakend
