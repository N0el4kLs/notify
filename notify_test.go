package notify

import (
	"testing"

	"notify/pkg/notifications/bark"
	"notify/pkg/notifications/wechat"
)

const (
	BARK_SERVER string = "http://10x.xxx.xxx.xxx"
	DEVICE_KEY  string = "UMuZzk*******"

	APP_ID      string = "wx**********"
	APP_SERCRET string = "3156ecdf5e**********"
	WECHATID    string = "oAdJS5_HB***********"
	TEMPLATEID  string = "Zzt_-MGXR******"
)

func TestNotice(t *testing.T) {
	message := bark.InitMessage("Hello, this is bark message")
	conf := bark.Config{
		BarkServer: BARK_SERVER,
		DeviceKey:  DEVICE_KEY,
	}
	barkProvider := bark.Init(conf, message)

	wechatTemplate := make(map[string]interface{})
	wechatTemplate["content"] = "Hello, this is wechat message1"
	wechatConf := wechat.Config{
		AppID:      APP_ID,
		AppSecret:  APP_SERCRET,
		WechatID:   WECHATID,
		TemplateID: TEMPLATEID,
	}
	wechatProvider := wechat.Init(wechatConf, wechatTemplate)

	Notice(barkProvider, wechatProvider)
}
