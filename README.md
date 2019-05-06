## 工程简介
本工程分两部分，src/main为用户合约的目录，src/xchain为合约底层运行机制的目录

## 包结构简介
src/main用于存放用户合约，目前只支持单个文件的合约，目录固定为src/main/go/contract

src/xchain用于存放合约底层运行所需的所有文件

* src/xchain/main.go为入口，负责初始化和启动整套机制
* src/xchain/base为基础，负责启动服务端和客户端
* src/xchain/client为连接Peer的客户端，为src/xchain/stub提供支持
* src/xchain/stub为客户端的桩，为用户合约提供支持
* src/xchain/service为客户端的接口定义，指明用户合约需要提供的接口
* src/xchain/proto为合约相关*.proto生成的*.pb.go
* src/xchain/model为所有的数据模型定义

## 使用步骤简介
1.将src/xchain和src/Dockerfile拷贝到拥有Docker的环境中，比如~/workdir

2.打包基础镜像并导出成tar包，便于拷贝到xchain所在平台
```
$ cd ~/workdir
$ sudo docker build -t xchain-contract-goenv:3.6 .
$ sudo docker save -o goenv.tar xchain-contract-goenv:3.6
```

3.在本工程中编写用户合约，编译无异常后拷贝src/main/go/contract/user_contract.go，使用工具上传合约



