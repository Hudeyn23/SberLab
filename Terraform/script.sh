#!/bin/bash

apt update
apt ugrade -y
apt install -y golang
mkdir /root/SberLab
git clone https://github.com/borodun/SberLab.git /root/SberLab/
cd /root/SberLab/Server/ || exit
go build
./Server
