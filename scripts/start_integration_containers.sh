#!/bin/sh

docker rm --force integration_db_server

docker run -d -p 9042:9042 -p 9142:9142 -v $PWD/schema:/data --name integration_db_server cassandra:2.2

sleep 30

docker exec -it integration_db_server cqlsh -f /data/init.cql 
