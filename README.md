# :coin: Charlie

![](https://img.shields.io/github/v/release/amirhnajafiz/charlie)
![](https://img.shields.io/badge/platform-docker-9cf)
![](https://img.shields.io/badge/language-python-blue)
![](https://img.shields.io/badge/language-bash-success)
![](https://img.shields.io/badge/env-okd4-yellow)

Docker image for executing python code from ```gitlab/github``` repository.
You can use this image to execute python scripts on ```okd``` or any other platform without
needing any pipeline, dockerfile or anything.
Just login with your ```snapp-cloud``` account into the namespace that you want.
Create an image registery, build and push this image into it and use it.

In ```snappline``` we use ```charlie``` in order to execute ```python``` scripts as ```crob jobs``` in ```okd4```.

## :unlock: Login

```sh
docker login -u <okd4-user> -p <okd4-token> <namespace>
```

## :heavy_check_mark: Requirements

Make sure to have these files:

- ```script.py``` file which contains the main fuction of the script.
- ```requirements.txt``` project dependencies.

Also make sure that your ```gitlab/github``` repository is public or accessible from your namespace.

## :hammer: Build

First build your image:

```shell
docker build . --build-arg REPOSITORY='https://gitlab.snapp.ir/snappline/api.git' --build-arg SCRIPT_PATH='script/report' --build-arg DIRECTORY='api' -t <snapp-image-registery>/<namespace>/charlie:v0.1.0 -f build/Dockerfile
```

### :wrench: Build Arguments

- ```REPOSITORY``` git repository url. Example: ```https://gitlab.snapp.ir/snappline/api.git```
- ```DIRECTORY``` name of the directory when it's done cloning. Example: ```api```
- ```SCRIPT_PATH``` directory of script. Example: ```script/report```

## :pushpin: Push

Push image to your namespace ```image stream```:

```shell
docker push <snapp-image-registery>/<namespace>/charlie:v0.1.0
```

## :bomb: Run

### :whale: Docker

Use the built image in order to execute your code:

```shell
docker run <snapp-image-registery>/charlie@v0.1.0
```

### :ship: Kubernetes Job

```yml
apiVersion: batch/v1
kind: CronJob
metadata:
  name: charlie
  labels:
    app: kubernetes/charlie
    type: cronjob/charlie
spec:
  schedule: "* * * * *"
  jobTemplate:
    spec:
      template:
        spec:
          containers:
          - name: charlie-coffin
            image: <registery>/<namespace>/charlie:v0.1.0
            imagePullPolicy: IfNotPresent
          restartPolicy: Never
```
