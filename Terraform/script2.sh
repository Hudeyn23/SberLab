#!/bin/bash

apt update
apt upgrade -y
apt install docker.io
docker pull borodun/front
docker pull borodun/back
sudo docker run -p 80:5000 borodun/front
sudo docker run -p 9999:9999 borodun/back


