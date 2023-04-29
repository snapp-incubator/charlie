# from a python based image (bullseye)
FROM python:3.10-bullseye

LABEL maintainers='amirhnajafiz,mehditeymorian'
LABEL company='snapp'
LABEL team='snappline'
LABEL app='charlie'
LABEL year=2023
LABEL pyversion=3.10
LABEL core=bash

# set environement variables
ENV REPOSITORY=''
ENV SCRIPT_PATH=''
ENV DIRECTORY=''

# create a workdirectory
WORKDIR src/

# create clone directory
RUN mkdir clone

# copy the core script file
COPY core/script.sh script.sh

# add execute
RUN chmod +x ./script.sh

# start script
CMD ./script.sh