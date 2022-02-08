# Furbot-MiraiGo

移植于[Furbot-Mirai](https://github.com/furleywolf/Furbot-Mirai)的go开源版本

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

### 作为模块安装(基于MiraiGo-Template进行开发)

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
