#!/bin/sh
rm /home/charliepi/go/atsGoStrip/assets/index.html;
rm /home/charliepi/go/atsGoStrip/assets/*.js;
rm /home/charliepi/go/atsGoStrip/assets/index*.css;
rm -rf /home/charliepi/go/atsGoStrip/assets/images

cd /home/charliepi/astro/AlphaTreeService;
npm run build;
mv /home/charliepi/astro/AlphaTreeService/dist/index.html /home/charliepi/go/atsGoStrip/assets/;
mv /home/charliepi/astro/AlphaTreeService/dist/assets/*.js /home/charliepi/go/atsGoStrip/assets/;
mv /home/charliepi/astro/AlphaTreeService/dist/assets/*.css /home/charliepi/go/atsGoStrip/assets/;
mv /home/charliepi/astro/AlphaTreeService/dist/assets/images /home/charliepi/go/atsGoStrip/assets/;
# cd /home/charliepi/go/atsGoStrip;
# docker-compose up --build;