#!/bin/bash
APP_NAME=gotubebot

function with_docker {
  read -p "Create .env file? [y/n] " STATUS
  if [[ "$STATUS" = "y" ]]; then
    read -p "TOKEN=" TOKEN
    echo "TOKEN=$TOKEN" >> .env
    echo "POSTGRES=True" >> .env
    echo "POSTGRES_HOST=postgres" >> .env
  fi
  wget --no-check-certificate https://raw.githubusercontent.com/Devil666face/${APP_NAME}/main/docker-compose.yaml
}

function with_systemd {
  wget --no-check-certificate https://github.com/Devil666face/${APP_NAME}/releases/latest/download/${APP_NAME}.tgz && \
      tar -xf ${APP_NAME}.tgz && \
      rm -rf ${APP_NAME}.tgz

  read -p "Create systemd unit? [y/n] " STATUS
  if [[ "$STATUS" = "y" ]]; then
    echo "[Unit]" >> $APP_NAME.service
    echo "Description=$APP_NAME telegram" >> $APP_NAME.service
    echo "After=network.target" >> $APP_NAME.service
    echo "" >> $APP_NAME.service
    echo "[Service]" >> $APP_NAME.service
    echo "User=www-data" >> $APP_NAME.service
    echo "Group=www-data" >> $APP_NAME.service
    echo "WorkingDirectory=%PWD%" >> $APP_NAME.service
    echo "ExecStart=%PWD%/$APP_NAME" >> $APP_NAME.service
    echo "Restart=on-failure" >> $APP_NAME.service
    echo "" >> $APP_NAME.service
    echo "[Install]" >> $APP_NAME.service
    echo "WantedBy=multi-user.target" >> $APP_NAME.service

    sed -i 's|%PWD%|'"$PWD"'|g' ./$APP_NAME.service
    ln -s $PWD/$APP_NAME.service /etc/systemd/system/$APP_NAME.service
    systemctl daemon-reload
    systemctl enable $APP_NAME.service --now
  fi

  read -p "Create .env file? [y/n] " STATUS
  if [[ "$STATUS" = "y" ]]; then
    read -p "TOKEN=" TOKEN
    echo "TOKEN=$TOKEN" >> .env
  fi
  chown -R www-data:www-data ../$APP_NAME
  systemctl restart $APP_NAME.service
}

function main {
  mkdir -p $APP_NAME
  cd $APP_NAME

  read -p "Deploy use docker? [y/n] " USE_DOCKER
  if [[ "$USE_DOCKER" = "y" ]]; then
    with_docker
  else
    with_systemd
  fi
}

main