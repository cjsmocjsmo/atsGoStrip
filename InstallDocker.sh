#!/bin/sh

# curl -fsSL https://download.docker.com/linux/debian/gpg | gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg && \
# echo \
#   "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/debian \
#   $(lsb_release -cs) stable" | tee /etc/apt/sources.list.d/docker.list > /dev/null && \
# apt-get update && \
# apt-get dist-upgrade -y && \
# apt-get auto-clean -y && \
# apt-get auto-remove -y && \
# apt-get install -y docker-ce docker-ce-cli containerd.io

docker run docker/compose:1.29.2 version;
echo alias docker-compose="'"'docker run --rm \
    -v /var/run/docker.sock:/var/run/docker.sock \
    -v "$PWD:$PWD" \
    -w="$PWD" \
    docker/compose:1.29.2'"'" >> ~/.bashrc;
source ~/.bashrc;