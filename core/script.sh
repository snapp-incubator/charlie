#!/usr/bin/env bash

echo '[OK] cloning into' "$REPOSITORY"

cd clone || exit

echo '[OK] into clone'

# remove if exists
rm -rf "$DIRECTORY"

# clone into repo
git clone "$REPOSITORY"

echo '[OK] repo cloned!'

# get into repository in order to run the script
cd "$DIRECTORY"/"$SCRIPT_PATH" || exit

# install dependencies
pip install -r requirements.txt || exit

echo '[OK] install requirements!'

# execute the script
python script.py