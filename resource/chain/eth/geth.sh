#!/bin/bash
/usr/bin/geth --http --datadir ./node0 --dev --dev.period 1 --mine --miner.threads 2 \
    --http.api 'eth,net,web3,miner,personal' \
    --http.addr 0.0.0.0 \
    --allow-insecure-unlock > /var/log/geth.log 2>&1 &

# ./geth --http --datadir ./node0 --dev --dev.period 1 --mine --miner.threads 2 \
#     --http.api 'eth,net,web3,miner,personal' \
#     --http.addr 0.0.0.0 \
#     --allow-insecure-unlock > geth.log 2>&1 &