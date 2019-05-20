FROM node:8.16.0-alpine
LABEL MAINTAINER="Irakli Nadareishvili"

COPY . /apidocs
WORKDIR /apidocs

RUN apk upgrade --update \
 && apk add --no-cache build-base \
 && npm i -g redoc-cli

CMD redoc-cli bundle -o index.html api.yml