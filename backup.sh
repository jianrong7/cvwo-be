#!/bin/bash
docker exec -t cvwo-be-db-1 pg_dumpall -c -U jianrong > dump_$(date +%Y-%m-%d_%H_%M_%S).sql
mv /home/ec2-user/*.sql /home/ec2-user/backup