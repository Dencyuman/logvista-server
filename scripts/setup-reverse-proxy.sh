#!/bin/bash

# 安全のためにエラーが発生したらスクリプトを停止する
set -e

# server_nameをコマンドラインから入力
# パブリックipかドメイン名を入力してください
read -p "Enter your server_name (domain or public IP): " server_name

# Nginxのインストール
echo "Installing Nginx..."
brew install nginx

# Nginx設定ファイルの作成
NGINX_CONF="/home/linuxbrew/.linuxbrew/etc/nginx/servers/my_fastapi_app.conf"
echo "server {
    listen 80; # 外部からのアクセスを受けるポートを80に設定
    server_name $server_name;  # ユーザー入力を反映

    location / {
        proxy_pass http://localhost:8080; # 内部アプリケーションへの転送先を8080に変更
        proxy_set_header Host \$host;
        proxy_set_header X-Real-IP \$remote_addr;
        proxy_set_header X-Forwarded-For \$proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto \$scheme;
    }
}" | tee $NGINX_CONF

# Nginxの設定をテスト
echo "Testing Nginx configuration..."
nginx -t

# Nginxを再起動して設定を適用
echo "Restarting Nginx..."
brew services restart nginx

echo "Nginx has been configured successfully."
echo "You can access your FastAPI app at http://$server_name"
