#!/bin/bash

wait-for-it.sh db:42001 -- /opt/alert-service/service
