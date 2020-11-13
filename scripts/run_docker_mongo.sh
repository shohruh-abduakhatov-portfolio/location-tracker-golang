#!/bin/bash

docker pull mongo
docker container rm mongo
docker run --name api -p 127.0.0.1:27017:27017  -d mongo