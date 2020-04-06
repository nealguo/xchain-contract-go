## 工程简介
本工程src/xchain为合约底层运行机制的目录，src/data为grpc相关证书

## 包结构简介
src/xchain用于存放合约底层运行所需的所有文件

* src/xchain/main.go为入口，负责初始化和启动整套机制
* src/xchain/contract/starter.go负责启动服务端和客户端
* src/xchain/contract/stub.go为客户端的桩，为用户合约提供支持
* src/xchain/contract/client.go为连接Peer的客户端，为stub.go提供支持
* src/xchain/contract/service.go为客户端的接口定义，指明用户合约需要提供的接口
* src/xchain/proto为合约相关*.proto生成的*.pb.go
* src/xchain/model为所有的数据模型定义

## 使用步骤简介
1.将文件夹拷贝到拥有Docker的环境中，比如~/workdir
  * src/xchain
  * src/Dockerfile
  * src/data

2.打包基础镜像并导出成tar包，便于拷贝到xchain所在平台
```
$ cd ~/workdir
$ sudo docker build -t xchain-contract-goenv:3.8 .
$ sudo docker save -o goenv.tar xchain-contract-goenv:3.8
```


