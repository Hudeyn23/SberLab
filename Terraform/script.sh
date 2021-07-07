#!/bin/bash

apt update
apt upgrade -y
apt install -y golang
mkdir /root/SberLab
git clone https://github.com/borodun/SberLab.git /root/SberLab/
cd /root/SberLab/Server/
go build
/root/SberLab/Server/Server
