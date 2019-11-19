#!/bin/bash

DB_ADDRESS=localhost:3306
DB_USERNAME=root
DB_PASSWORD=12345678

go build || exit
./tasa-vietnam-api -dbAddress=$DB_ADDRESS -username=$DB_USERNAME -password=$DB_PASSWORD
