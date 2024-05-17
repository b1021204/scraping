FROM ubuntu:18.04

MAINTAINER meruneru

ENV IMAGE_NAME=MikanOS

RUN apt-get update
RUN apt-get install -y vim \
build-essential \
python3 \
git \
okteta \
tmux \
dosfstools 
RUN apt install -y llvm-7

#https://github.com/uchan-nos/mikanos-build
RUN cd $HOME && git init && git clone https://github.com/uchan-nos/mikanos-build.git osbook && cd $HOME 
RUN apt install -y ansible;
RUN cd osbook/devenv && ansible-playbook -K -i ansible_inventory ansible_provision.yml

CMD ["/bin/bash"]
