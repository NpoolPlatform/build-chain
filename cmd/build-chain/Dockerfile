FROM centos:centos7.9.2009

COPY build-chain /usr/local/bin/build-chain
COPY BuildChain.viper.yaml /etc/BuildChain/BuildChain.viper.yaml
ENV ENV_DATA_DIR=/data \
    ENV_LOG_DIR=/var/log \
    ENV_LOG_LEVEL=debug \
    ENV_ETH_ENDPOINT="" \
    ENV_INVERSTOR_KEY=""

RUN yum update -y
RUN yum install ca-certificates curl -y

WORKDIR /usr/local/bin/

CMD ./build-chain run