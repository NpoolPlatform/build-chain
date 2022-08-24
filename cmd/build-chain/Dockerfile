FROM uhub.service.ucloud.cn/entropypool/alpine:3.16

COPY build-chain /usr/local/bin/build-chain

ENV ENV_LOG_DIR=/var/log \
    ENV_LOG_LEVEL=debug

ENTRYPOINT ["build-chain"]
CMD ["help"]
