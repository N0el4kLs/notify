package bark

import (
	"encoding/json"
	"fmt"

	"github.com/imroc/req/v3"
	"github.com/projectdiscovery/gologger"
)

type Bark struct {
	Clients *Options
}

// Options Configuration options required to run bark
type Options struct {
	Config
	Message Message
}

// Config you can all the parameters by your application
type Config struct {
	BarkServer string // the server address of running bark
	DeviceKey  string // the device key which your tools or apps generated
}

// Message bark notification body settings
type Message struct {
	Body  string `json:"body" yaml:"body"`
	Title string `json:"title" yaml:"title"`
	Badge int    `json:"badge" yaml:"badge"`
	Icon  string `json:"icon" yaml:"icon"`
	Group string `json:"group" yaml:"group"`
	Url   string `json:"url" yaml:"url"`
	Sound string `json:"sound" yaml:"sound"`
	Key   string `json:"device_key" yaml:"key"`
}

// InitMessage Todo
// This function can just set message text only, in future will support set more options
func InitMessage(body string) Message {
	return Message{
		Body: body,
	}
}

func Init(config Config, message Message) Bark {
	options := &Options{
		Config:  config,
		Message: message,
	}

	return Bark{
		Clients: options,
	}
}

func (b Bark) Notice() error {
	b.Clients.Message.Key = b.Clients.DeviceKey
	barkBody, err := json.Marshal(b.Clients.Message)
	if err != nil {
		gologger.Warning().Label("Bark").Msgf("Unmarshal bark message error %s\n", err)
		return err
	}
	c := req.C().SetBaseURL(fmt.Sprintf("%s/push", b.Clients.BarkServer))
	resp := c.Post().SetBodyJsonBytes(barkBody).Do()
	if resp.Err != nil {
		gologger.Warning().Label("Bark").Msgf("Send bark notification error %s\n", resp.Err)
		return resp.Err
	}
	return nil
}
