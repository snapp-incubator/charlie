#!/usr/bin/env bash

echo '[OK] cloning into' "$REPOSITORY"

cd clone || echo '[Failed] mounting' && ecit

# clone into repo
git clone "$REPOSITORY"

# get into repository in order to run the script
cd "$DIRECTORY"/"$SCRIPT_PATH" || echo "[Failed] repo not found" && exit

# install dependencies
pip install -r requirements.txt || echo '[Failed] install requirements!'

echo '[OK] install requirements!'

# execute the script
python script.py