#/bin/sh

set -B # enable brace expansion
for i in {1..10}; do
  curl -o /dev/null --silent --write-out '%{http_code}\n' 'http://krakend.kong-poc.kidsloop.live/rate-limit'
done
