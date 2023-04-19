# from a python based image (bullseye)
FROM python:3.10-bullseye

LABEL maintainers='amirhnajafiz,mehditeymorian'
LABEL team='snappline'

LABEL app=runtime
LABEL pyversion=3.10
LABEL core=bash

# set environement variables
ARG REPOSITORY=''
ARG SCRIPT_PATH=''
ARG DIRECTORY=''

# create a workdirectory
WORKDIR src/

# clone into
RUN git clone $REPOSITORY

# enter into project directory
WORKDIR $DIRECTORY/$SCRIPT_PATH

# copy the core script file
COPY core/script.sh script.sh

# add execute
RUN chmod +x ./script.sh

# start script
CMD ./script.sh