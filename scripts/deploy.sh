#!/bin/bash

DIR=`dirname $(readlink -f $0)`
OLDPWD=`pwd`

cd $DIR/../

# AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY must be provided via env

ecs-cli configure --region us-west-1 --cluster default --compose-project-name-prefix ""
#ecs-cli up --keypair marketwatcher --capability-iam --size 1 --instance-type t2.medium
ecs-cli compose --file docker-compose.yml service up

cd $OLDPWD
