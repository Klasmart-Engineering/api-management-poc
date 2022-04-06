FROM golang:1.17-alpine AS builder

ARG ENV

RUN apk add make gcc musl-dev

WORKDIR /tmp/builder
RUN mkdir plugin
COPY plugin plugin
WORKDIR /tmp/builder/plugin

RUN go build -buildmode=plugin -o krakend-debugger.so .

FROM devopsfaith/krakend as krakend

USER root

RUN mkdir /etc/krakend/config
COPY config/ /etc/krakend/config/
COPY krakend.json /etc/krakend/krakend.json

ENV FC_ENABLE=1
ENV FC_SETTINGS=/etc/krakend/config/settings/${ENV:-dev}
ENV FC_PARTIALS=/etc/krakend/config/partials
ENV FC_TEMPLATES=/etc/krakend/config/templates

RUN mkdir -p /etc/krakend/plugins
COPY  --from=builder /tmp/builder/plugin/krakend-debugger.so /etc/krakend/plugins/krakend-debugger.so
RUN chmod +x /etc/krakend/plugins/*.so

USER krakend
