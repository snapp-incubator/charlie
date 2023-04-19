#!/usr/bin/env bash

echo '[OK] cloning into' "$REPOSITORY"

# install dependencies
pip install -r requirements.txt || echo '[Failed] install requirements!'

echo '[OK] install requirements!'

# execute the script
python script.py