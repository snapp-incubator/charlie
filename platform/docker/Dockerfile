# using the base image ubuntu
FROM ubuntu:22.04

# setting ubuntu defaults
RUN apt-get update && apt-get install -y locales && rm -rf /var/lib/apt/lists/* \
	&& localedef -i en_US -c -f UTF-8 -A /usr/share/locale/locale.alias en_US.UTF-8 \

ENV LANG en_US.utf8

# install essiantials
RUN apt install build-essential checkinstall \
    libreadline-gplv2-dev  libncursesw5-dev libssl-dev \
    libsqlite3-dev tk-dev libgdbm-dev libc6-dev libbz2-dev

# install python, golang, gcc
# gcc
RUN apt install gcc-12

# python
