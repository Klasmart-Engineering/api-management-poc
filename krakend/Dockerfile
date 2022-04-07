FROM golang:1.17-alpine AS builder

ARG ENV

RUN apk add make gcc musl-dev

WORKDIR /tmp/builder
RUN mkdir plugin
COPY plugin plugin
WORKDIR /tmp/builder/plugin

RUN go build -buildmode=plugin -o krakend-debugger.so .

FROM 246958245973.dkr.ecr.eu-west-2.amazonaws.com/krakend-enterprise:kidsloop as krakend

USER root

RUN mkdir /etc/krakend/config
COPY config/ /etc/krakend/config/
COPY krakend.json /etc/krakend/krakend.json

ENV FC_ENABLE=1
ENV FC_SETTINGS=/etc/krakend/config/settings/${ENV:-dev}
ENV FC_PARTIALS=/etc/krakend/config/partials
ENV FC_TEMPLATES=/etc/krakend/config/templates

COPY  --from=builder /tmp/builder/plugin/krakend-debugger.so /opt/krakend/plugins/krakend-debugger.so
RUN chmod +x /opt/krakend/plugins/*.so

USER krakend
