#!/bin/bash

apt update
apt ugrade -y
apt install -y golang
git clone https://github.com/borodun/SberLab.git
cd SberLab/Server/ || exit
go build
./Server
