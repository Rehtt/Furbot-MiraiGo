/**
 * @Author: dsreshiram@gmail.com
 * @Date: 2022/2/8 16:47
 */

package furbot

import (
	"fmt"
	"github.com/Rehtt/Furbot-MiraiGo/furbot/net"
	"testing"
)

func TestBotParser(t *testing.T) {
	net.FurBotApi.Init("", "")
	fmt.Println(botParser("来只毛"))
	fmt.Println(botParser("来只 "))
	fmt.Println(botParser("找毛图 123"))
}
