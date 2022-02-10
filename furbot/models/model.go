/**
 * @Author: dsreshiram@gmail.com
 * @Date: 2022/2/7 16:24
 */

package models

type ConfStruct struct {
	QQ             string `yaml:"qq"`
	AuthKey        string `yaml:"authKey"`
	ShowTail       bool   `yaml:"showTail"`
	ResponseFriend bool   `yaml:"responseFriend"`
	ResponseGroup  bool   `yaml:"responseGroup"`
	Host           string `yaml:"host,omitempty"` // Host切换
}
