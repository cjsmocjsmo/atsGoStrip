#!/bin/sh

sudo apt-get update;
sudo apt-get -y dist-upgrade;
sudo apt-get -y install golang docker-compose;
sudo apt-get -y autoclean;
sudo apt-get -y autoremove;


sudo snap install core;
sudo snap refresh core;
sudo snap install --classic docker;
sudo snap install --classic certbot;
sudo ln -s /snap/bin/certbot /usr/bin/certbot;
sudo certbot certonly --standalone --preferred-challenges http -d atsdo.xyz;
cp /etc/letsencrypt/archive/atsdo.xyz/*.pem /root/atsGoStrip/

    # --cert-path /root/data --key-path /root/data --fullchain-path /root/data --chain-path /root/data;

# certbot certonly --dns-google -d atsio.xyz;

# sudo ls /etc/letsencrypt/live/atsdo.xyz  gives you location of certs to use in program