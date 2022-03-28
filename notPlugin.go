//go:build !plugin

/**
 * @Author: dsreshiram@gmail.com
 * @Date: 2022/3/28 14:50
 */

package main

import (
	// 载入Furbot模块
	_ "github.com/Rehtt/Furbot-MiraiGo/furbot"
)

func loadPlugin() {}
