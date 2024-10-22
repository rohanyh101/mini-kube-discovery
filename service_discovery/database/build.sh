#!/bin/bash

PASS="nopass"
PORT=6379

docker build --build-arg PASS=$PASS -t redis_img .

docker run -d --rm --network service_discovery \
    --name redis_cs \
    -p $PORT:$PORT \
    redis_img

echo "container 'redis server' is now up and running..."