#!/bin/sh

sudo apt-get update;
sudo snap install core;
sudo snap refresh core;
sudo snap install --classic certbot;
sudo ln -s /snap/bin/certbot /usr/bin/certbot;
sudo certbot certonly --standalone -d atsio.xyz

# certbot certonly --dns-google -d atsio.xyz;