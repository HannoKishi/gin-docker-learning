FROM alpine:latest
MAINTAINER "enming"
ADD . /app
WORKDIR /app
EXPOSE 8081
CMD "./gin-docker-test"