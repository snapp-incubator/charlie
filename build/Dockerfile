# from a python based image (bullseye)
FROM python:3.10-bullseye

MAINTAINER 'amirhnajafiz'
MAINTAINER 'mehditeymorian'

LABEL app=runtime
LABEL pyversion=3.10

# set environement variables
ENV REPOSITORY_PATH=''
ENV SCRIPT_PATH=''

# create a workdirectory
WORKDIR src/

# copy the core script file
COPY core/script.sh script.sh

# add execute
RUN chmod +x ./script.sh

# start script
CMD ./script.sh