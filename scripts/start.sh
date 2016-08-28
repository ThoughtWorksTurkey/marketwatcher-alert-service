#!/bin/bash

wait-for-it.sh -s -t30 db:9042 -- /opt/alert-service/service
