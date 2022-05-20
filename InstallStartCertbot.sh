#!/bin/sh

sudo apt-get update;
sudo snap install core;
sudo snap refresh core;
sudo snap install --classic certbot;
sudo ln -s /snap/bin/certbot /usr/bin/certbot;
certbot certonly --dns-digitalocean -d atsio.xyz;