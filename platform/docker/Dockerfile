# using the base image ubuntu
FROM ubuntu:22.04

# setting ubuntu defaults
RUN apt-get update && apt-get install

# setting the language
ENV LANG en_US.utf8
ENV PLS go,python,js,c,cpp
ENV PLATFORM docker

# install essiantials
RUN apt install build-essential -y

# install systemd
RUN apt-get install systemd -y

# install nginx
RUN apt install nginx -y \
    && systemctl enable nginx

# exposing port 80 for nginx
EXPOSE 80

# install python, golang, gcc
# gcc
RUN apt install gcc-12 -y

# python
RUN add-apt-repository ppa:deadsnakes/ppa
RUN apt install python2.7 -y \
    && apt install python3.11 -y

# golang
RUN apt install golang-go -y

# node js
RUN apt install nodejs npm -y

# set labels
LABEL app=runtime
LABEL platform=docker

# set maintainer
MAINTAINER amirhossein21

# keeping container alive
ENTRYPOINT ["tail"]
CMD ["-f", "/dev/null"]
