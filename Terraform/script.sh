#!/bin/bash

apt update
apt ugrade -y
apt install -y golang
git clone https://github.com/borodun/SberLab.git
cd SberLab/ServerDemo/ || exit
go build
./Server
