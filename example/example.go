package main

import (
	"os"

	"github.com/N0el4kLs/notify"
	"github.com/N0el4kLs/notify/pkg/notifications/bark"
)

func main() {
	msg := "Your message here"
	conf := bark.Config{
		BarkServer: os.Getenv("BARK_SERVER"),
		DeviceKey:  os.Getenv("DEVICE_KEY"),
	}

	// 设置需要发送的消息,目前只支持文本内容的设置
	message := bark.InitMessage(
		msg,
		bark.WithTitle("This is title for bark notification"),
		bark.WithGroup("Test"),
		bark.WithSound("shake"),
		bark.WithIcon("https://particles.oss-cn-beijing.aliyuncs.com/img/github_bark_icon-1686929068926.png"),
	)
	// 根据前面的 conf以及 messgae 生成 Provider对象
	barkProvider := bark.Init(conf, message)

	// 使用 options 模式运行传入的Provider对象
	notify.Notice(barkProvider)
}
