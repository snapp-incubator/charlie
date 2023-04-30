#!/usr/bin/env bash


cd clone || exit

echo '[OK] into clone'

rm -rf "$DIRECTORY"

git_address=""

if [ "$HTTP_SECURE" == true ]
then
  git_address="https://";
else
  git_address="http://";
fi

if [ "$GIT_SECURE" == true ]
then
  git_address="${git_address}${GIT_USER}:${GIT_TOKEN}@"
fi

git_address="${git_address}${REPOSITORY}"

echo '[OK] cloning into' "$git_address"

git clone "$git_address"

echo '[OK] repo cloned!'

cd "$DIRECTORY"/"$SCRIPT_PATH" || exit

echo '[OK] now running!'

python script.py