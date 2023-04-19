#!/usr/bin/env bash

echo '[OK] cloning into' "$REPOSITORY_PATH"

# clone into repository
git clone "$REPOSITORY_PATH"

# get inside repository
cd "$REPOSITORY_PATH" || echo 'failed to clone into repository' && exit
cd "$SCRIPT_PATH" || echo 'failed to find script directory' && exit

# install dependencies
pip install -r requirements.txt || echo 'failed to install requirements'

echo '[OK] install requirements!'

# execute the script
python script.py