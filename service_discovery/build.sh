#!/bin/bash

rm -rf ./bin
PORT=5555

go build -o ./bin/discovery

docker kill discovery

docker build -t discovery_img .

docker run -d -it --rm --network service_discovery \
    -p $PORT:$PORT \
    --name discovery \
    -e REDIS_ADDR=redis_cs:6379 \
    -e REDIS_PASS=nopass \
    -e REDIS_DB=0 \
    discovery_img

echo "container 'redis discovery' is now up and running..."