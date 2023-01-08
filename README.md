# scalpel

## 环境搭建相关

0. 设置临时环境变量
```
set GO111MODULE=on 
set GOPROXY=https://goproxy.io
```

1. `go mod init github.com/iamzowe/scalpel`
2. `go mod tidy`
3. `go mod vendor`

`go mod vendor` 遇到问题：
`github.com/sirupsen/logrus: module github.com/sirupsen/logrus: Get https://proxy.golang.org/github.com/sirupsen/logrus/@v/list: dial tcp 216.58.200.241:443: i/o timeout` 执行 `export GOPROXY=https://goproxy.io`

`go mod` 完成后goland的import包变红：
在`Preference->Go->Go Modules` 勾选 `Enable Go Modules integration`

## 配置 git pull 不需输入密码

进入项目的工程目录，执行如下目录即可
`git config credential.helper store`

