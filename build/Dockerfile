# from a python based image (bullseye)
FROM python:3.10-bullseye

LABEL maintainers='amirhnajafiz,mehditeymorian'
LABEL company='Snapp!'
LABEL team='Snappline'
LABEL app='Charlie'
LABEL platform='gitlab'
LABEL year=2023
LABEL pyversion=3.10
LABEL core=bash

# set environement variables
# repository is the git repository address
ENV REPOSITORY=''
# the path of python script in repository
ENV SCRIPT_PATH=''
# directory name of project when it's done clonning
ENV DIRECTORY=''
# http/https
ENV HTTP_SECURE=true
# git security parameters
ENV GIT_SECURE=false
ENV GIT_USER=''
ENV GIT_TOKEN=''

# create a workdirectory
WORKDIR src/

# create clone directory
RUN mkdir clone

# copy requirements
COPY libs/requirements.txt requirements.txt

# install all dependencies
RUN pip install -r requirements.txt

# copy the core script file
COPY core/script.sh script.sh

# add execute
RUN chmod +x ./script.sh

# start script
CMD ./script.sh