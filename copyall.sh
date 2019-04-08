#!/usr/bin/env bash

export GOOS=linux
export CGO_ENABLED=0
cd accountservice; go get; go build -o accountservice-linux-amd64;echo built `pwd`;cd ..
cd healthchecker; go get; go build -o healthchecker-linux-amd64;echo built `pwd`; cd ..
export GOOS=windows
cp healthchecker/healthchecker-linux-amd64 accountservice/
docker build -t tony/accountservice accountservice/
#docker service rm accountservice
#docker service create --name=accountservice --replicas=1 -p=6767:6767 tony/accountservice

