#!/bin/bash

IP=$(hostname -I | awk '{print $1}')
PORT=6060

go build -ldflags="-X main.ip=$IP -X main.port=$PORT" -o ./bin/order

docker kill order

docker build -t order_img .

docker run -it --network service_discovery \
    --rm \
    --name order \
    -p $PORT:$PORT \
    order_img

echo "container 'order' is now up and running..."