#!/bin/bash

docker run --name=mysql-tasa -d -p 3306:3306 -e MYSQL_USER=remote -e MYSQL_PASSWORD=12345678 -e MYSQL_DATABASE=tasa -e MYSQL_ROOT_PASSWORD=12345678 mysql/mysql-server:8.0 