#!/bin/bash
set -e

npm config set registry https://registry.npmmirror.com

echo "[*] 设置 GPG 签名配置"

mkdir -p ~/.gnupg
chmod 700 ~/.gnupg

# 设置 GPG 配置文件
echo "use-agent" >> ~/.gnupg/gpg.conf
echo "pinentry-mode loopback" >> ~/.gnupg/gpg.conf
echo "allow-loopback-pinentry" >> ~/.gnupg/gpg-agent.conf

# git config --global gpg.program gpg
# git config --global commit.gpgsign true
# git config --global user.signingkey ABCDEF1234567890

# 启动 gpg-agent
gpgconf --kill gpg-agent || true
gpgconf --launch gpg-agent

# 设置终端交互 tty
echo "export GPG_TTY=\$(tty)" >> ~/.bashrc
export GPG_TTY=$(tty)

# GPG 预热，缓存密码
# echo "warming up gpg-agent..."
# echo "test" | gpg --clearsign --pinentry-mode loopback || true

echo "[*] 初始化完成。请手动配置 Git 的 user.signingkey（如果还未设置）"
