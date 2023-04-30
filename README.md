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
docker login -u <okd4-user> -p <okd4-token> <registery>/<namespace>
```

## :heavy_check_mark: Requirements

Make sure to have these files:

- ```script.py``` file which contains the main fuction of the script.
- ```deps/requirements.txt``` project dependencies when building the image (this will setup your image with requirements that you need).

Also make sure that your ```gitlab/github``` repository is public or accessible from your namespace.

## :hammer: Build

First build your image:

```shell
docker build . -t <snapp-image-registery>/<namespace>/charlie:v0.1.0 -f build/Dockerfile
```

## :pushpin: Push

Push image to your namespace ```image stream```:

```shell
docker push <snapp-image-registery>/<namespace>/charlie:v0.1.0
```

## :wrench: Environment Variables

- ```REPOSITORY``` git repository url. Example: ```https://gitlab.snapp.ir/snappline/api.git```
- ```DIRECTORY``` name of the directory when it's done cloning. Example: ```api```
- ```SCRIPT_PATH``` directory of script. Example: ```script/report```

## :bomb: Run

### :whale: Docker

Use the built image in order to execute your code:

```shell
docker run \
--env REPOSITORY='https://github.com/amirhnajafiz/charlie.git' \
--env DIRECTORY='charlie' \
--env SCRIPT_PATH='test' \
--mount type=bind,source="$(pwd)/clone",target=/src/clone \
<snapp-image-registery>/charlie@v0.1.0
```

### :ship: Kubernetes Job

Example job that executes once in a day to do something.

#### Configmap

```yml
apiVersion: v1
kind: ConfigMap
metadata:
  name: report-configs
data:
  REPOSITORY: 'gitlab_repository'
  DIRECTORY: 'export'
  SCRIPT_PATH: 'script'
```

#### Job

```yml
apiVersion: batch/v1
kind: CronJob
metadata:
  name: report-job
  labels:
    app.snappcloud.io/created-by: snappline
    app: report
spec:
  schedule: "* * */1 * *"
  jobTemplate:
    spec:
      template:
        spec:
          containers:
            - name: charlie
              image: <registery>/<namespace>/charlie:v0.3.0
              volumeMounts:
                - mountPath: /src/clone
                  name: report-dir
              envFrom:
                - configMapRef:
                    name: report-configs
          restartPolicy: Never
          volumes:
            - name: report-dir
              emptyDir: { }
```
