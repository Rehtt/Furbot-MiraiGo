package main

import (
	"flag"
	"github.com/Logiase/MiraiGo-Template/bot"
	"github.com/Logiase/MiraiGo-Template/utils"
	"github.com/Mrs4s/MiraiGo/client"
	"os"
	"os/signal"
	// 载入Furbot模块
	//_ "github.com/Rehtt/Furbot-MiraiGo/furbot"
)

func init() {
	utils.WriteLogToFS(utils.LogInfoLevel, utils.LogWithStack)
	//config.Init()
}

func main() {
	flag.Parse()

	f, err := os.Open("./device.json")
	if err != nil {
		bot.GenRandomDevice()
	} else {
		f.Close()
	}

	// 加载插件
	loadPlugin()

	// 初始化，使用二维码登录
	bot.Instance = &bot.Bot{}
	bot.Instance.QQClient = client.NewClientEmpty()

	// 使用协议
	// 不同协议可能会有部分功能无法使用
	// 在登陆前切换协议
	bot.UseProtocol(bot.AndroidPhone)

	// 初始化 Modules
	bot.StartService()

	// 登录
	bot.Login()

	// 保存登录状态
	bot.SaveToken()

	// 刷新好友列表，群列表
	bot.RefreshList()

	// 监听中断信号
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	// 优雅退出
	bot.Stop()
}
