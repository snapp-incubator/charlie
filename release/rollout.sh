#!/usr/bin/env bash

echo "Building $1 ..."

# make a docker release
docker build . -f platform/docker/Dockerfile -t amirhossein21/runtime:"$1"