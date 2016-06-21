#!/bin/bash

wait-for-it.sh db:9042 -- /opt/alert-service/service
