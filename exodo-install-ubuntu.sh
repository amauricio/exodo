#!/usr/bin/env bash

version=$(lsb_release --release | cut -f2)
codename=$(lsb_release --codename | cut -f2)
machine=$(uname -m)
nodename=$(uname -n)


echo $nodename
echo 'Instalando exodo...'
echo $version

if [ "$EUID" -ne 0 ]
  then echo "Run as root bitch!!!"
  exit
fi

##adding repository
add-apt-repository "deb http://apt.postgresql.org/pub/repos/apt/ $version-pgdg main"


##downloading key
wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | sudo apt-key add -

##update repository
apt update

##installing requirements
apt -y install build-essential libssl-dev  postgresql-9.6 golang-go git

##cloning repository from google git
echo "Get golang git from Google"
mkdir -p /usr/local/exodo/golang
git clone https://go.googlesource.com/go /usr/local/exodo/golang
git --git-dir /usr/local/exodo/golang/.git checkout go1.8.3
##installing golang
echo "Installing golang 1.8.3"
./all.bash
