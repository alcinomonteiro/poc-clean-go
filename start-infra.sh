#!/bin/bash

docker-compose -f "docker-compose-infra.yaml" up -d

sleep 10

docker exec mongodb /scripts/replicaSet-init.sh