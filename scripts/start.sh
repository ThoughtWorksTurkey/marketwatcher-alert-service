#!/bin/bash

wait-for-it.sh -s -t 30 db:9042 -- /opt/alert-service/service
