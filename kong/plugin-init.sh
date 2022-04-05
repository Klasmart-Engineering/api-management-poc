kubectl set env deployment/kong-istio-kong KONG_PLUGINS="bundled, simple-plugin" --containers="proxy" -n kong-istio
kubectl set env deployment/kong-istio-kong KONG_PLUGINSERVER_NAMES="simple-plugin" --containers="proxy" -n kong-istio
kubectl set env deployment/kong-istio-kong KONG_PLUGINSERVER_SIMPLE-PLUGIN_QUERY_CMD="/usr/local/bin/go-plugins/simple-plugin -dump" --containers="proxy" -n kong-istio

