//go:build plugin

/**
 * @Author: dsreshiram@gmail.com
 * @Date: 2022/3/28 14:48
 */

package main

import (
	"flag"
	"fmt"
	"log"
	"plugin"
)

type pluginPath []string

func (p *pluginPath) String() string {
	return fmt.Sprint(*p)
}
func (p *pluginPath) Set(value string) error {
	*p = append(*p, value)
	return nil
}

var plugins pluginPath

func init() {
	flag.Var(&plugins, "p", "插件地址")
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
