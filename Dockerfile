FROM ubuntu:20.04

ENV APP_NAME "gotubebot"
ENV TOKEN ""
ENV DEPS "wget"

RUN DEBIAN_FRONTEND=noninteractive \
    apt-get update --quiet --quiet && \
    apt-get upgrade --quiet --quiet && \
    apt-get install --quiet --quiet --yes \
    --no-install-recommends --no-install-suggests \
    ${DEPS} \
    && apt-get --quiet --quiet clean \
    && rm --recursive --force /var/lib/apt/lists/* /tmp/* /var/tmp/*

WORKDIR /var/www/${APP_NAME}

RUN wget --no-check-certificate https://github.com/Devil666face/${APP_NAME}/releases/latest/download/${APP_NAME}.tgz && \
    tar -xf ${APP_NAME}.tgz && \
    rm -rf ${APP_NAME}.tgz

CMD ["/bin/bash","-c","./$APP_NAME"]