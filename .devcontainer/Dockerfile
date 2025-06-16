ARG NODE_VERSION
FROM node:${NODE_VERSION}-slim

ENV PNPM_HOME=/home/node/.local/share/pnpm
ENV PATH=$PNPM_HOME:$PATH

RUN sed -i 's/deb.debian.org/mirrors.ustc.edu.cn/g' /etc/apt/sources.list.d/debian.sources \
    && apt-get update && apt-get upgrade -y \
    && apt-get install -y vim curl git gnupg2 gnupg-agent \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

RUN npm config set registry https://registry.npmmirror.com \
    && npm install -g pnpm && mkdir -p $PNPM_HOME && chown -R node:node $PNPM_HOME
RUN pnpm config set registry https://registry.npmmirror.com \
    && mkdir -p /home/node/.ssh /home/node/.gnupg && chown -R node:node /home/node
