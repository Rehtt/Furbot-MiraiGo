/**
 * @Author: dsreshiram@gmail.com
 * @Date: 2022/2/7 15:52
 */

package furbot

import (
	"github.com/Logiase/MiraiGo-Template/bot"
	"github.com/Logiase/MiraiGo-Template/utils"
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
	"github.com/Rehtt/Furbot-MiraiGo/furbot/models"
	FurBotApi "github.com/Rehtt/Furbot-MiraiGo/furbot/net"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"sync"
)

type ar struct{}

var instance = &ar{}
var logger = utils.GetModuleLogger("rehtt/furbot-miraiGo")
var conf *models.ConfStruct

func init() {
	bot.RegisterModule(instance)
}

func (a *ar) MiraiGoModule() bot.ModuleInfo {
	return bot.ModuleInfo{
		ID:       "rehtt/furbot-miraiGo",
		Instance: instance,
	}
}

func (a *ar) Init() {
	fileData, err := ioutil.ReadFile("./furbot.yaml")
	if err != nil {
		genConfig()
	}
	err = yaml.Unmarshal(fileData, &conf)
	if err != nil {
		logger.WithError(err).Error("'furbot.yaml' parsing failed")
		os.Exit(1)
	}
	// 初始化
	FurBotApi.Init(conf.QQ, conf.AuthKey)
}

func (a *ar) PostInit() {
}

func (a *ar) Serve(bot *bot.Bot) {
	if conf.ResponseGroup {
		// 群消息
		bot.OnGroupMessage(func(client *client.QQClient, m *message.GroupMessage) {
			match, text, img := botParser(m.ToString())
			if !match {
				return
			}
			msg := message.NewSendingMessage().Append(message.NewText(text))
			if img != nil {
				image, err := client.UploadGroupImage(m.GroupCode, img)
				if err != nil {
					logger.WithError(err).Error("图片上传失败")
					return
				}
				msg.Append(image)
			}
			client.SendGroupMessage(m.GroupCode, msg, true)
		})
	}
	if conf.ResponseFriend {
		// 私聊
		bot.OnPrivateMessage(func(client *client.QQClient, m *message.PrivateMessage) {
			match, text, img := botParser(m.ToString())
			if !match {
				return
			}
			msg := message.NewSendingMessage().Append(message.NewText(text))
			if img != nil {
				image, err := client.UploadPrivateImage(m.Sender.Uin, img)
				if err != nil {
					logger.WithError(err).Error("图片上传失败")
					return
				}
				msg.Append(image)
			}

			client.SendPrivateMessage(m.Sender.Uin, msg)
		})
	}
}

func (a *ar) Start(bot *bot.Bot) {
}

func (a *ar) Stop(bot *bot.Bot, wg *sync.WaitGroup) {
	defer wg.Done()
}
