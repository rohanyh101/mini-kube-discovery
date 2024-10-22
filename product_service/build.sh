#!/bin/bash

IP=$(hostname -I | awk '{print $1}')
PORT=7070

go build -ldflags="-X main.ip=$IP -X main.port=$PORT" -o ./bin/product

docker kill product

docker build -t product_img .

docker run -it --network service_discovery \
    --rm \
    --name product \
    -p $PORT:$PORT \
    product_img

echo "container 'product' is now up and running..."