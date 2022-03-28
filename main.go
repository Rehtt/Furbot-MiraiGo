package main

import (
	"flag"
	"fmt"
	"github.com/Logiase/MiraiGo-Template/bot"
	"github.com/Logiase/MiraiGo-Template/utils"
	"github.com/Mrs4s/MiraiGo/client"
	"log"
	"os"
	"os/signal"
	"plugin"
	// 载入Furbot模块
	//_ "github.com/Rehtt/Furbot-MiraiGo/furbot"
	//_ "github.com/Rehtt/Furbot-MiraiGo/send"
)

func init() {
	utils.WriteLogToFS(utils.LogInfoLevel, utils.LogWithStack)
	//config.Init()
}

type pluginPath []string

func (p *pluginPath) String() string {
	return fmt.Sprint(*p)
}
func (p *pluginPath) Set(value string) error {
	*p = append(*p, value)
	return nil
}

var plugins pluginPath

func main() {
	flag.Var(&plugins, "p", "插件地址")
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

func loadPlugin() {
	if len(plugins) == 0 {
		return
	}
	for _, p := range plugins {
		_, err := plugin.Open(p)
		if err != nil {
			log.Printf("%s 插件加载失败：%v", p, err)
		}
	}
}
