#!/bin/bash

docker build -t go-websocket-tutorial .
docker run -it -p 8080:8080 go-websocket-tutorial
