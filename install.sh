#!/bin/bash

apt-get update

apt-get install snap -y
snap install --channel=1.20 go --classic

echo 'SERVICENAME="stagetest"'>> /etc/environment

apt-get install nginx -y
ufw allow 'Nginx Full'
ufw allow 'OpenSSH'
yes |  ufw enable

 rm /etc/nginx/sites-available/default

echo 'server {
  # listen on port 80 (http)
  listen 80;
  server_name stagetest.leanafywms.com;
  
  location / {
    # redirect any requests to the same URL but on https
    proxy_pass http://localhost:8000;
    proxy_redirect off;
    proxy_set_header Host $host;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
  }
}' >> /etc/nginx/sites-available/default


 systemctl reload nginx.service
 systemctl restart nginx.service


echo '[Unit]
Description=stagetest

[Service]
Type=simple
LimitNOFILE=1024
Restart=on-failure
RestartSec=5s
WorkingDirectory=/home/ubuntu/test
ExecStart=/home/ubuntu/test/main
PermissionsStartOnly=true
StandardOutput=syslog
StandardError=syslog

[Install]
WantedBy=multi-user.target' >> /lib/systemd/system/stagetest.service

 chmod 755 /lib/systemd/system/stagetest.service





