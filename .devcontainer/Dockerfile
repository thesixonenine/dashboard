ARG NODE_VERSION=22.16.0-slim
FROM node:${NODE_VERSION}

RUN sed -i 's/deb.debian.org/mirrors.ustc.edu.cn/g' /etc/apt/sources.list.d/debian.sources \
    && apt-get update && apt-get upgrade -y \
    && apt-get install -y --no-install-recommends \
       vim curl git gnupg2 gnupg-agent openssh-client ncat pinentry-curses sudo \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/* \
    # add sudoers
    && echo node ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/node \
    && chmod 0440 /etc/sudoers.d/node \
    # set registry
    && npm config set registry https://registry.npmmirror.com
