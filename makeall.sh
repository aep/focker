#!/bin/sh

for i in *.docker
do
    n=aaep/focker-$(basename $i .docker)
    docker build -f $i -t $n .
done
