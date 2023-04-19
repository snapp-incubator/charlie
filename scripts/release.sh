#!/usr/bin/env bash

docker build . \
    --build-arg REPOSITORY='$1' \
    --build-arg SCRIPT_PATH='$2' \
    --build-arg DIRECTORY='$3' \
    -t "$4"/charlie:"$5" \
    -f build/Dockerfile
