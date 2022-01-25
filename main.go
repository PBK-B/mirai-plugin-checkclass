/*
 * @Author: Bin
 * @Date: 2021-10-02
 * @FilePath: /mirai-plugin-checkclass/main.go
 */
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"sync"

	"github.com/Logiase/MiraiGo-Template/bot"
	"github.com/Logiase/MiraiGo-Template/utils"
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
)

func main() {}

func RegisterModule() {

	bot.UnModule("bin.checkclass")

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

		mKeys := []string{"搜题 ", "搜题"}
		for _, value := range mKeys {

			// 正则匹配题目
			flysnowRegexp := regexp.MustCompile(value + `(.+)$`)
			params := flysnowRegexp.FindStringSubmatch(msg.ToString())
			if len(params) <= 0 {
				// 判断没有匹配到关键词，进入下一次循环
				continue
			}
			qKey := params[1] // 提取匹配到的关键词
			// logger.Infof("搜索题目=%+v\n", params[1])

			if qKey != "" {

				out, err := seek(qKey)
				if out == "" || err != nil {
					logger.Infof("搜索错误: %+v", err.Error())
					return
				}

				// 生成回复消息并发送
				m := message.NewSendingMessage().Append(message.NewText(out))
				msgid := c.SendPrivateMessage(msg.Sender.Uin, m)
				logger.Infof("回复: %+v", msgid.Id)
				// fmt.Printf("回复=%+v\n", msgt.Id)
				// 匹配成功一次之后就跳出匹配
				return
			}
		}

	})

}

func (mov *checkclass) Start(bot *bot.Bot) {
}

func (mov *checkclass) Stop(bot *bot.Bot, wg *sync.WaitGroup) {
	defer wg.Done()
}

// 插件额外功能实现函数
type Question struct {
	Type    int64
	Content string
	Answer  string
}

type Callback struct {
	Code int64
	Msg  string
	Data Question
}

func seek(q string) (callback string, err error) {
	seekApi := "https://study.jszkk.com/api/open/seek?q=" + q
	str, err := httpGet(seekApi)
	if err != nil {
		return "", err
		// t.Fatal(err)
	}
	back := &Callback{}
	if err = json.Unmarshal([]byte(str), back); err != nil {
		return "", err
	}

	callback = "题目：" + back.Data.Content + "\n答案：" + back.Data.Answer + "\n答案来源: 全能搜题 https://so.jszkk.com"
	return
}

func httpGet(url string) (res string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body), err
}
