# Charlie

![](https://img.shields.io/github/v/release/amirhnajafiz/charlie)
![](https://img.shields.io/badge/platform-docker-9cf)
![](https://img.shields.io/badge/language-python-blue)

Runtime for executing python code. You can use this image to execute python scripts on ```okd``` without
needing any pipeline or anything. Just login with your ```snapp-cloud``` account into the namespace that you want.
Create an image registery and push this image into it.

## Requirements

- ```script.py``` file which contains the main fuction of the script.
- ```requirements.txt``` project dependencies.

## Build

First build your image:

```shell
docker build . --build-arg REPOSITORY='https://github.com/amirhnajafiz/runtime' --build-arg SCRIPT_PATH='test' --build-arg DIRECTORY='runtime' -t amirhossein21/charlie:v0.1.5 -f build/Dockerfile
```

### Arguments

- ```REPOSITORY``` git repository url
- ```DIRECTORY``` name of the directory when it's done cloning
- ```SCRIPT_PATH``` directory of script

Use the following image in order to execute your code:

```shell
docker run amirhossein21/charlie@latest
```
