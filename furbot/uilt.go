/**
 * @Author: dsreshiram@gmail.com
 * @Date: 2022/2/8 10:58
 */

package furbot

import (
	"github.com/Rehtt/Furbot-MiraiGo/furbot/models"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

// 创建配置文件
func genConfig() {
	out, _ := yaml.Marshal(models.ConfStruct{})
	err := ioutil.WriteFile("./furbot.yaml", out, 0644)
	if err != nil {
		logger.WithError(err).Error("文件写入失败")
		os.Exit(1)
	}
	logger.Println("模板已生成，请修改furbot.yaml文件再运行")
	os.Exit(0)
}
