package feishu

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/imroc/req/v3"
	"github.com/projectdiscovery/gologger"
)

const (
	FEI_SHU_BASE_URL = "https://open.feishu.cn/open-apis/bot/v2/hook/%s"
)

type FeiShu struct {
	Clients *Options
}

// Options Configuration options required to run feishu
type Options struct {
	Config
	Message Message
}

// Config  You can get all the parameters in
// https://open.feishu.cn/document/client-docs/bot-v3/add-custom-bot
type Config struct {
	HookToken string // HookToken in bot information
	CryptoKey string // CryptoKey in bot information
}

type Message struct {
	Timestamp string `json:"timestamp"`
	Sign      string `json:"sign"`
	MsgType   string `json:"msg_type"`
	Content   struct {
		Text string `json:"text"`
	} `json:"content"`
}

// InitMessage feishu notification body settings
// current support text message only
func InitMessage(cont string) Message {
	return Message{
		MsgType: "text",
		Content: struct {
			Text string `json:"text"`
		}{
			Text: cont,
		},
	}
}

// Init create feishu instance
func Init(config Config, message Message) FeiShu {
	options := &Options{
		Config:  config,
		Message: message,
	}

	return FeiShu{
		Clients: options,
	}
}

func (f FeiShu) Notice() error {
	if f.Clients.HookToken == "" || f.Clients.CryptoKey == "" {
		gologger.Error().Label("FeiShu").Msgf("FeiShu HookToken or CryptoKey is empty\n")
		return errors.New("FeiShu HookToken or CryptoKey is empty\n")
	}
	sign, timestamp, err := genSign(f.Clients.CryptoKey)
	if err != nil {
		gologger.Error().Label("FeiShu").Msgf("Generate FeiShu sign error %s\n", err)
		return err
	}
	f.Clients.Message.Sign = sign
	f.Clients.Message.Timestamp = fmt.Sprintf("%v", timestamp)
	c := req.C().SetBaseURL(fmt.Sprintf(FEI_SHU_BASE_URL, f.Clients.HookToken))
	feishuBody, err := json.Marshal(f.Clients.Message)
	if err != nil {
		gologger.Warning().Label("FeiShu").Msgf("Unmarshal FeiShu message error %s\n", err)
		return err
	}
	resp := c.Post().SetBodyJsonBytes(feishuBody).Do()
	if resp.Err != nil {
		gologger.Warning().Label("FeiShu").Msgf("Send FeiShu notification error %s\n", resp.Err)
		return resp.Err
	}
	return nil
}

// genSign generate sign for feishu
func genSign(secret string) (string, int64, error) {
	timestamp := time.Now().Unix()
	//timestamp + key 做sha256, 再进行base64 encode
	stringToSign := fmt.Sprintf("%v", timestamp) + "\n" + secret
	var data []byte
	h := hmac.New(sha256.New, []byte(stringToSign))
	_, err := h.Write(data)
	if err != nil {
		return "", -1, err
	}
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return signature, timestamp, nil
}
