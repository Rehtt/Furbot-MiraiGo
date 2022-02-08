/**
 * @Author: dsreshiram@gmail.com
 * @Date: 2022/2/8 12:02
 */

package furbot

import (
	"bytes"
	"fmt"
	"github.com/Rehtt/Furbot-MiraiGo/furbot/constants"
	"github.com/Rehtt/Furbot-MiraiGo/furbot/models"
	"github.com/Rehtt/Furbot-MiraiGo/furbot/net"
	"io"
	"strings"
)

// 解析消息
func botParser(str string) (match bool, text string, img io.ReadSeeker) {
	var data models.ApiResponse
	var err error
	defer func() {
		if err != nil {
			logger.WithError(err).Error("解析失败")
		}
	}()
	if str == constants.FurPrefixRandom {
		data, err = net.FurBotApi.GetFurRandom()
		if err == nil {
			match = true
			text = fmt.Sprintf(
				"---每日吸毛Bot---\n"+
					"今天你吸毛了嘛？\n"+
					"FurID: %d\n"+
					"毛毛名字: %s\n"+
					"搜索方法:全局随机", data.Id, data.Name)
			img = bytes.NewReader(data.Image)
		}
	} else if strings.HasPrefix(str, constants.FurPrefixGetName) {
		name := strings.Join(strings.Split(str, constants.FurPrefixGetName), "")
		data, err = net.FurBotApi.GetFurByName(name)
		if err == nil {
			match = true
			text = fmt.Sprintf(
				"---每日吸毛Bot---\n"+
					"FurID: %d\n"+
					"毛毛名字: %s\n"+
					"搜索方法:模糊搜索", data.Id, data.Name)
			img = bytes.NewReader(data.Image)
		}
	} else if strings.HasPrefix(str, constants.FurPrefixGetId) {
		id := strings.Join(strings.Split(str, constants.FurPrefixGetId), "")
		data, err = net.FurBotApi.GetFurByID(id)
		if err == nil {
			match = true
			text = fmt.Sprintf(
				"---每日吸毛Bot---\n"+
					"FurID: %d\n"+
					"毛毛名字: %s\n"+
					"搜索方法:按FurID搜索", data.Id, data.Name)
			img = bytes.NewReader(data.Image)
		}
	}
	return
}
