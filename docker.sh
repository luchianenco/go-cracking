#!/usr/bin/env bash

docker-compose build
docker-compose up

##log into the container
docker-compose stop
