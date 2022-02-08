/**
 * @Author: dsreshiram@gmail.com
 * @Date: 2022/2/8 16:41
 */

package net

import (
	"fmt"
	"github.com/Rehtt/Furbot-MiraiGo/furbot/constants"
	"strconv"
	"testing"
	"time"
)

func TestFurBotApi_GetFurRandom(t *testing.T) {
	FurBotApi.Init("", "")
	data, err := FurBotApi.GetFurRandom()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(data)
}

func TestFurBotApi_GetFurByName(t *testing.T) {
	FurBotApi.Init("", "")
	data, err := FurBotApi.GetFurByName("")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(data)
}

func TestFurBotApi_GetFurByID(t *testing.T) {
	FurBotApi.Init("", "")
	data, err := FurBotApi.GetFurByID("")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(data)
}

func TestSign(t *testing.T) {
	FurBotApi.Init("", "")
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	fmt.Println(FurBotApi.sign(constants.ApiGetFursuitRand, timestamp))
	fmt.Println(FurBotApi.sign(constants.ApiGetFursuitByName, timestamp))
	fmt.Println(FurBotApi.sign(constants.ApiGetFursuitById, timestamp))
}
