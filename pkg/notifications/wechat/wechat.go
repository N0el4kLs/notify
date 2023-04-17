package wechat

import (
	"encoding/json"
	"fmt"

	"github.com/imroc/req/v3"
	"github.com/projectdiscovery/gologger"
)

const (
	WECHAT_PUBLIC_PLATFORM_URL = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
	WECHAT_TEMPLATE            = "https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=%s"
)

type WeChat struct {
	Client *Options
}

// Options Configuration options required to run wechat
type Options struct {
	Config
	Message Message
}

// Config  You can get all the parameters in
// https://mp.weixin.qq.com/debug/cgi-bin/sandbox?t=sandbox/login
type Config struct {
	AppID      string // appID in test account information
	AppSecret  string // appsecret in test account information
	WechatID   string // you can get WeChat id by scanning the test two-dimension code
	TemplateID string //  you can get the template id by create a template
}

// Message WeChat notification body settings
type Message struct {
	ToUser     string                 `json:"touser"`
	TemplateID string                 `json:"template_id"`
	Data       map[string]interface{} `json:"data"`
}

func Init(config Config, m map[string]interface{}) WeChat {
	var message map[string]interface{}
	message = make(map[string]interface{}, 10)
	for key, value := range m {
		message[key] = map[string]interface{}{
			"value": value,
		}
	}

	options := &Options{
		Config: config,
		Message: Message{
			Data: message,
		},
	}

	return WeChat{
		Client: options,
	}
}

func (w WeChat) Notice() error {
	// First Step
	// Get the WeChat api active token
	wechatAPI_GetToken := fmt.Sprintf(WECHAT_PUBLIC_PLATFORM_URL, w.Client.AppID, w.Client.AppSecret)
	c := req.C().SetBaseURL(wechatAPI_GetToken)
	response := c.Get().Do()
	if response.Err != nil {
		gologger.Warning().Label("Wechat").Msgf("Send wechat message error %s\n", response.Err)
		return response.Err
	}
	var respData resp
	err := response.Into(&respData)
	if err != nil {
		gologger.Warning().Label("Wechat").Msgf("wechat get access token error %s\n", err)
		return err
	}

	// Second Step
	// Using the active token you just get and send the message using specified template
	wechatAPI_Send := fmt.Sprintf(WECHAT_TEMPLATE, respData.AccessToken)
	w.Client.Message.ToUser = w.Client.Config.WechatID
	w.Client.Message.TemplateID = w.Client.Config.TemplateID

	messageBody, err := json.Marshal(w.Client.Message)
	if err != nil {
		gologger.Warning().Label("WeChat").Msgf("Unmarshal wechat message error %s\n", err)
		return err
	}
	c2 := req.C().SetBaseURL(wechatAPI_Send)
	r2 := c2.Post().SetBodyJsonBytes(messageBody).Do()
	if r2.Err != nil {
		gologger.Warning().Label("WeChat").Msgf("Send wechat message error %s\n", r2.Err)
		return r2.Err
	}
	return nil
}

type resp struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}
