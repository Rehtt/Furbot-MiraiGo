# Furbot-MiraiGo
[![Go Report Card](https://goreportcard.com/badge/github.com/Rehtt/Furbot-MiraiGo)](https://goreportcard.com/report/github.com/Rehtt/Furbot-MiraiGo)

移植于[Furbot-Mirai](https://github.com/furleywolf/Furbot-Mirai)的go开源版本

基于[MiraiGo-Template](https://github.com/Logiase/MiraiGo-Template)开发

## 安装

### 直接安装

```shell
go install github.com/Rehtt/Furbot-MiraiGo@latest
Furbot-MiraiGo
```

OR

```shell
git clone https://github.com/Rehtt/Furbot-MiraiGo
cd Furbot-MiraiGo
go install
Furbot-MiraiGo
```
OR

```shell
git clone https://github.com/Rehtt/Furbot-MiraiGo
cd Furbot-MiraiGo
go build .
./Furbot-MiraiGo
```

### 作为模块安装(基于MiraiGo-Template二次开发)

在适当位置引入
```go
package main
import (
	//...
    _ "github.com/Rehtt/Furbot-MiraiGo/furbot"
	//...
)
```
参考[main.go](./main.go)

### 作为插件(golang plugin)
> 使用插件模式需要使用cgo，只支持linux，并且不能交叉编译
> 
> 因为插件需要重新打包，所以总体积会比模块模式大，但更灵活

执行命令：
```shell
go build -tags plugin -o Furbot-MiraiGo .
cd furbot/pluginMode
go build -buildmode=plugin -o furbot.so .
cd -
./Furbot-MiraiGo -p furbot/pluginMode/furbot.so
```

## 使用
第一次运行时会自动生成`furbot.yaml`文件，修改后再次运行即可
文件内容详解：
```yaml
#此处填写申请开源版地址的QQ号码
qq: 0

#此处填写申请开源版地址的授权码
authKey: 0

#是否响应私聊消息 true 响应  false 不响应
responseFriend: true

#是否响应群聊消息
responseGroup: true
```
