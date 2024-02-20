#!/bin/bash

# 手動部分
#sudo yum update -y
#sudo yum install git -y
#git clone https://github.com/Dencyuman/logvista-server.git

# NVMを使ってNode.jsをインストール
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.1/install.sh | bash
export NVM_DIR="$HOME/.nvm"
[ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"
[ -s "$NVM_DIR/bash_completion" ] && \. "$NVM_DIR/bash_completion"
nvm install v16 # 最新バージョンのNode.jsをインストール
nvm use v16

# Go言語のインストール
GO_VERSION="1.21.1" # 使用したいGoのバージョンに応じて変更してください
wget https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz
echo "export PATH=\$PATH:/usr/local/go/bin" >> ~/.bash_profile
source ~/.bash_profile

# インストールの確認
node -v
npm -v
go version
