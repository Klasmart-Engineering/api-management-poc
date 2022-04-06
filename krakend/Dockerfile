FROM devopsfaith/krakend
COPY krakend.json /etc/krakend/krakend.json

USER root
RUN mkdir -p /etc/krakend/plugins
COPY plugin/*.so /etc/krakend/plugins
RUN chmod +x /etc/krakend/plugins/*.so

USER krakend
