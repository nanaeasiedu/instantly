#!/bin/bash

git pull

if sudo docker stop web-instant; then sudo docker rm web-instant fi;
sudo docker rmi web-instant
sudo docker build -t web-instant .
sudo docker run -d -p 3000:3000 --name web-instant web-instant
