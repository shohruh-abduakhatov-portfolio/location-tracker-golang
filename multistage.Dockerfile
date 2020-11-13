FROM golang:1.10-alpine3.8 AS multistage

LABEL authors="abduakhatov,Adilbek11kz" \
    maintainer="abduakhatov,Adilbek11kz" \
    version="1.0" \
    description="ascii-asrt-web for Alem school. All rights reserved."

RUN mkdir /app
ADD . /app/
WORKDIR /app/ascii-art-web/app
RUN go build

## MULTISTAGE TO RUN BUILT

FROM alpine:3.8
RUN mkdir /app
COPY --from=multistage /app/ /app/

WORKDIR /app/ascii-art-web/app
EXPOSE 8080:8080
CMD ["./app"]