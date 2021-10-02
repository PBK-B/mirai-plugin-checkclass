/*
 * @Author: Bin
 * @Date: 2021-10-02
 * @FilePath: /bot_checkclass/main.go
 */
package main

import (
	"fmt"
	"sync"

	"github.com/Logiase/MiraiGo-Template/bot"
	"github.com/Logiase/MiraiGo-Template/utils"
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
)

func main() {}

func RegisterModule() {
	bot.RegisterModule(checkclassInstance)

	bot.StartService()

	fmt.Printf("[安装模块]: %s 成功!\n", checkclassInstance.MiraiGoModule().ID)
}

var checkclassInstance = &checkclass{}

var logger = utils.GetModuleLogger("bin.checkclass")

// var tem map[string]string

type checkclass struct {
}

func (mov *checkclass) MiraiGoModule() bot.ModuleInfo {
	return bot.ModuleInfo{
		ID:       "bin.checkclass",
		Instance: checkclassInstance,
	}
}

func (mov *checkclass) Init() {
	fmt.Printf("【Bot Plugins】模块初始化=%+v\n", mov.MiraiGoModule().ID)
}

func (mov *checkclass) PostInit() {
}

func (mov *checkclass) Serve(b *bot.Bot) {

	b.OnPrivateMessage(func(c *client.QQClient, msg *message.PrivateMessage) {

		// fmt.Printf("message=%+v\n", msg.ToString())

		if botObj := bot.Instance; botObj == nil {
			// 机器人已下线，直接结束回复流程
			fmt.Println("【收到消息】机器人已下线，直接结束回复流程")
			return
		}

		if msg.ToString() == "/help" {
			m := message.NewSendingMessage().Append(message.NewText(`【指令菜单】
  /help (帮助菜单)
  /info (系统状态)
  `))
			c.SendPrivateMessage(msg.Sender.Uin, m)
			return
		}

		if msg.ToString() == "/info" {
			m := message.NewSendingMessage().Append(message.NewText("系统状态: 正常\n机器人状态: 在线\n"))
			c.SendPrivateMessage(msg.Sender.Uin, m)
			return
		}

	})

}

func (mov *checkclass) Start(bot *bot.Bot) {
}

func (mov *checkclass) Stop(bot *bot.Bot, wg *sync.WaitGroup) {
	defer wg.Done()
}
