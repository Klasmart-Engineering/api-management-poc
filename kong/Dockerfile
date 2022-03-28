FROM kong/kong-gateway:2.8.0.0-alpine
USER root
# Example for GO:
COPY ./plugins/simple-plugin/simple-plugin /usr/local/bin/go-plugins/simple-plugin

USER kong
# ENTRYPOINT ["/docker-entrypoint.sh"]

# EXPOSE 8000 8443 8001 8444
# STOPSIGNAL SIGQUIT
# HEALTHCHECK --interval=10s --timeout=10s --retries=10 CMD kong health

# CMD ["kong", "docker-start"]
