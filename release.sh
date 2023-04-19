#!/usr/bin/env bash

docker build . --build-arg REPOSITORY='https://github.com/amirhnajafiz/runtime' --build-arg SCRIPT_PATH='test' --build-arg DIRECTORY='runtime' -t "$1"/runtime:"$2" -f build/Dockerfile
