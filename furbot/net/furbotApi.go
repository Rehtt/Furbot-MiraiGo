/**
 * @Author: dsreshiram@gmail.com
 * @Date: 2022/2/8 15:01
 */

package net

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/Rehtt/Furbot-MiraiGo/furbot/constants"
	"github.com/Rehtt/Furbot-MiraiGo/furbot/models"
	"io"
	"strconv"
	"time"
)

type furBotApi struct {
	q       string
	authKey string
	init    bool
}

var FurBotApi *furBotApi

func (f furBotApi) GetFurRandom() (out models.ApiResponse, err error) {
	return f.getFurBasic(constants.ApiGetFursuitRand, nil)
}
func (f furBotApi) GetFurByID(fid string) (out models.ApiResponse, err error) {
	return f.getFurBasic(constants.ApiGetFursuitById, map[string]string{"fid": fid})
}
func (f furBotApi) GetFurByName(name string) (out models.ApiResponse, err error) {
	return f.getFurBasic(constants.ApiGetFursuitByName, map[string]string{"name": name})
}

func (f furBotApi) getFurBasic(api string, extraQuery map[string]string) (out models.ApiResponse, err error) {
	if !f.init {
		return models.ApiResponse{}, fmt.Errorf("Furbot未初始化")
	}
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	query := map[string]string{
		"qq":        f.q,
		"timestamp": timestamp,
		"sign":      f.sign(api, timestamp),
	}
	for k, v := range extraQuery {
		query[k] = v
	}
	response, err := httpGet(constants.ApiHost+api, nil, query)
	if err != nil {
		return models.ApiResponse{}, err
	}
	json.Unmarshal(response, &out)

	image, err := httpGet(out.Url, nil, nil)
	if err != nil {
		return models.ApiResponse{}, err
	}
	if len(image) == 0 {
		return models.ApiResponse{}, fmt.Errorf("图像拉取失败")
	}
	out.Image = image
	return
}

func Init(q, authKey string) {
	FurBotApi = &furBotApi{authKey: authKey, q: q, init: true}
}

// 签名
func (f furBotApi) sign(apiPath, timestamp string) (sign string) {
	w := md5.New()
	io.WriteString(w, fmt.Sprintf("%s-%s-%s", apiPath, timestamp, f.authKey))
	return fmt.Sprintf("%x", w.Sum(nil))
}
