#!/bin/bash

where=$PWD
if [ "$EUID" -ne 0 ]
  then echo "Run as root bitch!!!"
  exit
fi

version=$(lsb_release --release | cut -f2)
codename=$(lsb_release --codename | cut -f2)
machine=$(uname -m)
nodename=$(uname -n)


echo $nodename
echo 'Instalando exodo...'
echo $version

##adding repository
add-apt-repository "deb http://apt.postgresql.org/pub/repos/apt/ $codename-pgdg main"
apt-add-repository 'deb https://apt.dockerproject.org/repo ubuntu-xenial main'

##add key
sudo apt-key adv --keyserver hkp://p80.pool.sks-keyservers.net:80 --recv-keys 58118E89F3A912897C070ADBF76221572C52609D

##downloading key
wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | sudo apt-key add -

##update repository
apt update

##installing requirements
apt -y install curl build-essential libssl-dev  postgresql-9.6 git docker-engine python-pip

##instal docker compose
pip install docker-compose

##cloning repository from google git
echo "Get golang"

#creating exodo folder
mkdir -p /usr/local/exodo/golang
cd /usr/local/exodo/golang

#download source golang
curl -O https://storage.googleapis.com/golang/go1.8.linux-amd64.tar.gz

#unpack golang
sha256sum go1.8.linux-amd64.tar.gz
tar xvf go1.8.linux-amd64.tar.gz

##permissions
sudo chown -R root:root ./go

#moving source golang 
sudo mv go /usr/local

#environment
echo "GOPATH=$HOME/work" >> ~/.profile
echo "PATH=$PATH:/usr/local/go/bin:$GOPATH/bin" >> ~/.profile
export GOROOT=$HOME/go
export GOPATH=$HOME/work
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin

#return main folder
cd $PWD
