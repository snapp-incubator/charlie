# runtime

Runtime for executing python code.

## Requirements

- ```script.py``` file which contains the main fuction of the script.
- ```requirements.txt``` project dependencies.

## Build

Use the following image in order to execute your code:

```sh
docker run -d -e REPOSITORY_PATH='gitlab.com/group/repository' -e SCRIPT_PATH='scripts/report' amirhossein21/runtime@latest
```

The following command will clone into ```REPOSITORY_PATH``` and executes the script inside ```SCRIPT_PATH```.
