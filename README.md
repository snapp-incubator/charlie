<h1 align="center">
    Runtime
</h1>

<div align="center">
    <img src="https://img.shields.io/github/v/release/amirhnajafiz/runtime?display_name=tag&style=flat-square" />
    <img src="https://img.shields.io/badge/target-Docker-blue?style=flat-square" />
</div>

<br />

This is a runtime image in order to create
containers for executing you codes 
inside it. With this image, you can execute
your codes inside an isolated environment, which
gives you an advantage for debugging your codes.
This image can be used for DOM Judge, 
Code execution platforms, etc.

## Docker Image

Get the latest docker image of runtime ```docker pull amirhossein21/runtime:latest```.

### Kubernetes

You can set the runtime on kubernetes by using the same image.

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: runtime
  labels:
    app: runtime
spec:
  containers:
    - name: runtime
      image: amirhossein21/runtime:latest
      ports:
        - containerPort: 80
```

### Programming Languages Support

- ```C```
- ```C++```
- ```Node.js```
- ```Python v2.7```
- ```Python v3.11```
- ```Golang v1.18```
