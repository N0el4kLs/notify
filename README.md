<h2 align='center'>Notify</h2>

## 介绍

自用的一款消息通知模块,打算将常见的第三方消息通知整合到一起,并作为Go的第三方包整合到其他项目中.后续考虑会增加命令行运行模式

目前已集成的第三方通知:

- [x] ~~Bark(IOS)~~

- [x] ~~Wechat~~
- [x] Feishu
- [ ] Enterprise WeChat
- [ ] DingDing



## 使用

使用前,请先将改包导入到您的项目中:

```
go get github.com/N0el4kLs/notify
```



具体相关通知配置如下:

### Bark 配置

```go
// 初始化bark的配置信息,包括 bark server 的地址以及使用bark应用程序生成的 devicekey
conf := bark.Config{
	BarkServer: BARK_SERVER,
	DeviceKey:  DEVICE_KEY,
}

// 设置需要发送的消息,目前只支持文本内容的设置
message := bark.InitMessage(
	"Hello, this is bark message"
	WithTitle("This is title for bark notification"),
	WithGroup("Test"),
	WithSound("shake"),
	WithIcon("https://particles.oss-cn-beijing.aliyuncs.com/img/github_bark_icon-1686929068926.png"),
	)

// 根据前面的 conf以及 messgae 生成 Provider对象
barkProvider := bark.Init(conf, message)

// 使用 options 模式运行传入的Provider对象
Notice(barkProvider)
```



### 微信公众平台配置

```go
// 前往微信公众平台获取 appid, appsecret, wechatid以及tempalteid参数
// https://mp.weixin.qq.com/debug/cgi-bin/sandboxinfo?action=showinfo&t=sandbox/
// 具体获取方式请自行百度
wechatConf := wechat.Config{
	AppID:      APP_ID,
	AppSecret:  APP_SERCRET,
	WechatID:   WECHATID,
	TemplateID: TEMPLATEID,
}

// 在创建消息模板时,可能使用了参数形式,
// 举例,我在创建模板的正文中填写的内容为  这个消息的类型为{{contype.DATA}},他的正文如下:{{content.DATA}}
// 你可以在Golang中通过创建 map 类型,以键key为模板的变量名, 值value为对应变量名填充的值
// 例如:
wechatTemplate := make(map[string]interface{})
wechatTemplate["contype"] = "Hello, this is wechat message"
wechatTemplate["content"] = "Content: HelloWorld"

// 根据前面的 config 以及 wechatTemplate 生成 wechat Provider 对象
wechatProvider := wechat.Init(wechatConf, wechatTemplate)

// 使用 options 模式运行传入的Provider对象
Notice( wechatProvider)
```

### 飞书配置

```go
// 设置飞书机器人的配置信息,包括飞书机器人的 HookToken以及CryptoKey
feishuConf := feishu.Config{
		HookToken: HOOK_TOKEN,
		CryptoKey: CRYPTO_KEY,
	}
	
// 设置需要发送的消息,目前只支持文本内容的设置	
feishuMessage := feishu.InitMessage("Hello, this is feishu message")

// 创建飞书的Provider对象
feishuProvider := feishu.Init(feishuConf, feishuMessage)

// 使用 options 模式运行传入的Provider对象
notify.Notice(feishuProvider)
```



## 日志

### 2023-4-18

目前只支持 `Bark(IOS)` 以及 `微信公众号(WeChat)` 两种消息通知方式.

Todo：

- [ ] `Bark`消息初始化函数优化
- [ ] 增加常用消息模板,使其满足大都数场景
- [ ] 报错提示优化


### 2023-6-13

新增支持 `飞书(Feishu)` 消息通知方式

### 2023-6-16

新增 `Bark` 通知设置，包括: `Title`, `Group`, `Sound`, `Icon`.