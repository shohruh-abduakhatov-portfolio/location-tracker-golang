#!/bin/bash

docker container run -p 9090:8181 --detach --name gtservice gt


# image_name="ascii-art-web-docker"

# docker build -f Dockerfile.multistage -t $image_name .
# docker run --name ascii-art-web -p 8080:8080 --rm -d $image_name