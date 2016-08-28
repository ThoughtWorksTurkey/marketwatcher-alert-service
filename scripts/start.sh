#!/bin/bash

wait-for-it.sh -s db:9042 -- /opt/alert-service/service
