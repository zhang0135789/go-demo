# go-demo

eg 
  用go语言实现的简单区块链demo https://github.com/BlockChain-dragon/blockchain-demo

go语言交流qq群：786937587




### demo22
    由于example_cc.go中引入了fabric的依赖，导致项目引用依赖报错，因此引入go mod去解决引入依赖错误的问题
    初始化go mode
        go mod init example_cc.go
        go get github.com/hyperledger/fabric/core/chaincode/shim@v1.4
        go get github.com/hyperledger/fabric/protos/peer@v1.4
    执行以上命令，将解决依赖缺失问题。

#### golang在1.11版本后引入了go mode 模块来管理go的依赖，有点类似与java的maven

