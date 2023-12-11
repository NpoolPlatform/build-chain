# Npool Build-Chain

[![Test](https://github.com/NpoolPlatform/build-chain/actions/workflows/main.yml/badge.svg?branch=master)](https://github.com/NpoolPlatform/build-chain/actions/workflows/main.yml)

[目录](#目录)

- [Npool Build-Chain](#npool-build-chain)
  - [说明](#说明)
    - [新增功能](#新增功能)
  - [功能](#功能)
  - [命令](#命令)
  - [最佳实践](#最佳实践)
  - [环境变量](#环境变量)
  - [使用说明](#使用说明)
    - [1 部署build-chain server](#1-部署build-chain-server)
    - [2 扒取合约信息(此步骤可省略，deploy/erc20.csv有拉取好的合约信息)](#2-扒取合约信息此步骤可省略deployerc20csv有拉取好的合约信息)
    - [3 部署合约](#3-部署合约)
    - [sphinx-plugin支持](#sphinx-plugin支持)
  - [补充](#补充)
    - [ethereum测试链启动指南](#ethereum测试链启动指南)
  - [增加合约币流程](#增加合约币流程)

-----------

## 说明

![功能展示](docs/smart_coin.png)
目前仅支持ethereum上的部分erc20币种（top100中支持19个币种）

### 新增功能

- [x] eth链交互
- [x] 快捷erc20最基础合约部署
- [ ] 部署erc20非proxy合约
- [ ] 部署erc20 proxy合约
- [x] 统一存入地址的大小写
- [ ] build-chain server可以不和eth-wallet部署到同一个物理环境
- [x] 基础erc20接水
- [x] eth接水
- [x] 水龙头功能
- [ ] 快捷erc20非proxy合约部署
- [ ] tron链交互
- [ ] filecoin链交互
- [ ] 自动显示账户、币种细节（web页面选中自动查询币种信息、账户信息）
- [ ] 流程上支持快速加入一个币种

## 功能

合约工具服务

- [x] 单独作为工具服务存在
- [x] 可用做cli配合已启动的服务工作

## 命令

- make init ```初始化仓库，创建go.mod```
- make verify ```验证开发环境与构建环境，检查code conduct```
- make verify-build ```编译目标```
- make test ```单元测试```
- make generate-docker-images ```生成docker镜像```
- make build-chain ```单独编译服务```
- make build-chain-image ```单独生成服务镜像```
- make deploy-to-k8s-cluster ```部署到k8s集群```

## 最佳实践

- 每个服务只提供单一可执行文件，有利于 docker 镜像打包与 k8s 部署管理
- 每个服务提供 http 调试接口，通过 curl 获取调试信息
- 集群内服务间 direct call 调用通过服务发现获取目标地址进行调用

## 环境变量

| 变量名称             | 支持的值                            | 说明       |
|:---------------- |:------------------------------- |:-------- |
| ENV_ETH_ENDPOINT | ip:port default(127.0.0.1:8545) | 必填参数-用于server |
| ENENV_INVERSTOR_KEY | privateKey | 可选参数-用于部署合约和提供水龙头手续费（请提供有一定ETH的私钥） |

## 使用说明

### 1 部署build-chain server

1.1 依赖coinbase部署

需要在钱包机所在的环境启动server端，
启动时需要配置eth测试链endpoint地址，
还需要运行build-chain目录下有BuildChain.viper.yaml文件

```Shell
# 默认获取本地 eth 钱包
./build-chain run

# 指定私钥部署
./build-chain run --ee http://EthereumWalletIP:Port --ik $wallet_privatekey
```

1.2 docker部署
```shell
# docker启动
docker run --name buildchain --restart=always -p 50491:50491 -p 50490:50490 -v /opt/chain/buildchain/data:/data -e ENV_ETH_ENDPOINT="eth_wallet_endpoint" -e ENENV_INVERSTOR_KEY="privateKey" buildchain:test

# G网环境启动示例
docker run --name buildchain --restart=always -p 50491:50491 -p 50490:50490 -v /opt/chain/buildchain/data:/data -e ENV_ETH_ENDPOINT="https://goerli.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161" -e ENE_INVERSTOR_KEY="7a87e4528e013e533d63dd7661ead74fc3b25289652469a289bdf89b84e15c21" buildchain:test
```

注：数据目录在/data，按需挂载到物理机或传入ENV_DATA_DIR环境变量修改

### 2 扒取合约信息(此步骤可省略，deploy/erc20.csv有拉取好的合约信息)

使用crawl爬取ethscan.io上的合约数据，在真正使用时请用deploy/erc20.csv下的合约信息

```Shell
# 必要时需要设置代理
# export all_proxy=socks5://IP:PORT
# export https_proxy=socks5://IP:PORT
# export no_proxy='my_ip/netmask'

# 查看参数
./build-chain crawl

# 拉取ethscan中erc20-top-tokens 前1-100，大约有19个币种可以成功拉取
./build-chain crawl -o 1 -l 100

# 指定公网contract进行拉取
./build-chain crawl --co 0xdAC17F958D2ee523a2206206994597C13D831ec7
```

### 3 部署合约
使用事先准备好的合约信息（一般在deploy下）部署token，部署前要确保build-chain server所使用的钱包地址有币

例如：
```Shell
./build-chain deploy --file ./deploy/erc20.csv --bc-server ServerIP:50491

# 清除旧合约，重新部署
./build-chain deploy --file ./deploy/erc20.csv --bc-server ServerIP:50491 --force true
```

部署成功的合约可访问server端的web界面申请空投

在浏览器可访问服务页面

http://ServerIP:50490/

页面上可申请已经部署且可支持的合约币

### sphinx-plugin支持
sphinx-plugin在测试模式下启动需要build-chain server也启动（当前就eth测试环境需要），启动后查询需要的测试合约地址，当需要的合约类型不存在时会请求build-chain部署好并自动获取合约地址。

## 补充

### ethereum测试链启动指南

启动eth测试链时，需要指定http地址以及开启允许unlock

建议启动方式如下

```Shell
nohup geth --http --datadir ~/eth_node0 --dev --dev.period 1 --mine --miner.threads 2 --http.api 'eth,net,web3,miner,personal' --http.addr 0.0.0.0 --allow-insecure-unlock > ~/eth_node0/geth.log 2>&1 &

geth attach ~/eth_node0/geth.ipc
```

## 增加合约币流程

1. 启动build-chain server端
```
./build-chain run --eth-endpoint 'http://172.16.3.90:8545' --ik e86bead113c6500b7c9ef662ea6029ed71dcbe6a7e36e5c09945cdc3e616f788
```

2. 打开server的web页面（浏览器访问ServerIP:50490）,可查看已有币种信息

3. 使用build-chain cli工具将合约部署到测试网
```
./build-chain deploy --file ./deploy/erc20.csv --bc-server ServerIP:50491
```
