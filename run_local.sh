#!/bin/bash
#run local docker-hosted app

DB_ADDRESS=host.docker.internal:3306 #this address is for macos docker only ^^ | use docker network=host and 'localhost' for Linux distros
DB_USERNAME=root
DB_PASSWORD=12345678

docker build -t tasa_local . || exit
docker stop tasa
docker rm tasa
docker run -p 8000:8000 --name=tasa  tasa_local  ./tasa-vietnam-api -dbAddress=$DB_ADDRESS -username=$DB_USERNAME -password=$DB_PASSWORD
