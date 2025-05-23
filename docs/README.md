# VioGo

纯go编写的机器人业务后端，bot取名为 **Vio** ，qq机器人的实现基于gocq提供的api，基于gocq的上报事件调用转发外部讯息并返回。

可以完成基于chatgpt聊天任务，以及一些简单的指令响应。

**目前正在重构，并且开发后台管理界面！除了聊天,VioGo的目标是可以在后台分析聊天数据的奇妙bot！**

**如果你有想法和意见请提issue！这对我和bot都很重要！**
**欢迎qq加群讨论：340961300**

实现平台以及实现的外部功能:

- [X] QQ bot
  - [X] 随机涩图
  - [X] 猎杀对决枪声语音
  - [X] 发送已撤回的消息
  - [X] 禁言抽奖
- [ ] Telegram bot(暂时已停止开发)

## 写在前面/Preface

有部署聊天机器人的想法，但是我使用的国内服务器，而且服务器性能也一般，于是决定不用云服务器部署了，找个国外的Paas平台，把写的后端送上去就好了。

目前已经完成了：

- 天然支持http请求，使用go原生net包，创建了一个 `/post`请求的路由，可以解析post内容转发调用chatgpt的api
- ~~为Tg设置了webhook，可以监听tg服务器的消息，实现tgbot。针对Telegram的消息处理，对私人，群组，超级群组各有不同的应答模式。~~
- 配合[go-cqhttp](https://github.com/Mrs4s/go-cqhttp)使用，用于实现qqbot,基于上报事件，调用gocq提供的api进行回复。注意参考[gocq官方文档](https://docs.go-cqhttp.org/reference/#websocket)。可以调用api发送图片，随机图片api[参考](https://api.lolicon.app/setu/v2)

AI调用：

- 目前只支持了chatGPT，调用其api进行回复。该功能使用了[go-openai](https://github.com/sashabaranov/go-openai)库，支持了go语言对openai api的调用
- 提示词参考：[awesome-chatgpt-prompts-zh](https://github.com/PlexPt/awesome-chatgpt-prompts-zh/blob/main/prompts-zh.json)

## 结构说明

- AI： 提供ai服务，从ai_server文件中驱动，子文件为具体的api的调用实现。
- conf： 定义环境变量结构体
- cron: 使用cron包实现定时任务
- docs: 文档页面，使用 `docsify`
- gocq： qqbot的业务层核心。
  - command: 最业务层的逻辑，使用各种基础设施和其他服务，完成机器人指令的调用
  - cqCode: 定义qq消息中cq码的结构体，实现编码和解码
  - cqEvent: 属于gocq中事件的定义，只用于参考，未使用其结构体
  - event: 监听上报事情后的处理入口
  - gocq_resp_params: 响应参数的结构体定义
  - gocq_send_params: 发送参数的结构体定义
  - gocq_sender: gocq的发送者实例，用于api调用
  - gocq_server: 在这里注册了一个gocq的**单例**，全局调用gocq相关的基础设施
- server： 创建or使用服务(http,ws,tgbot)
  - handler: 处理http和请求以及ws请求
  - server: 使用依赖注入启动各种服务
- unused: 未实现，或者搁置的功能，如tgbot
- utils: 调用外界功能的函数库和工具库

## 实现要点

- 服务层注册基础设施，如redis(目前只搭载了redis缓存，用于存放部分聊天数据)，通过依赖注入的方式激活其他只能单个存在的服务，如gocq服务(依赖ws连接存在，ws连接只有一次，所以是个全局单例)。
- gocq服务在handler中进行ws连接的时候完全创建，因为ws连接读消息是阻塞的，为了实现并发，在gocq服务中，我创建了一个消息发送者，一个消息队列，以及一个响应map。handler中只处理接受到的上报事件，把接受到的响应存放在一个channel中，在具体需要响应数据的地方去pop。响应map是根据响应中的echo映射指定的响应数据。这样实现了消息接受，发送和处理的解耦，可以同时并发处理。
- redis的作用：最简单的场景，存放撤回的消息id，因为未知需要取用该撤回的消息的时机，存放在redis中方便需要的时候进行api调用取到该消息。目前使用的是heroku的add-ons的redis组件，也是通过环境变量读取redis的url地址进行注册：`REDISCLOUD_URL`
- AI服务: 预计支持多种ai的调用，目前已经完成了ai服务的代码解耦并且结构化，便于扩充。
- 前端设计: 预计做成一个后台管理面板，精力有限，会使用直接模板，注重功能的实现，不只聊天！
- 结构体设计完全遵循gocq文档
- 使用go的websocket库创建和qqbot的ws连接，为基于gocq的qqbot的服务入口。
- gocq的配置建议阅读我的[个人博客](http://viogami.tech/index.php/blog/144/)
- chatgpt的调用参考go的openai库文档即可，也很完善。注意调用api是无法进行联系上下文对话的，要实现上下文对话只有把历史消息都post给api，这显然是不现实的。

## 部署建议

我推荐放在paas上，key和密钥是通过环境变量读取的，也就是config中定义的字段。
因为tgbot要求后端webhook地址必须有证书，也就是https访问。一般paas平台都会自带证书的，放自己服务器上需要自己配证书服务，麻烦点，并且需要修改项目中读取环境变量的代码，改为读取配置文件(新建).

**我是放在heroku这个paas平台，国外的，可以通过github学生包免费使用，但是需要绑定国外visa卡。**

之前也有使用过Zeabur，国人开发的，discord上官方也很活跃，如果服务请求不多，可以使用，每月有5刀的免费额度，无需绑卡。

### paas推荐

- Heroku
- Zeabur

 更多paas相关信息，可以查询 [Free for dev](https://github.com/ripienaar/free-for-dev)中的列表.

Zeabur挺不错，国内社区，discord上回复也很即使，一键部署挺快的，github集成。
zeabur上项目部署非常快,甚至不用写dockfile,而且对go项目有完整的支持,算是符合他们的口号:

> Deploying your service with one click

但是Zeabur每月5$，如果是聊天机器人，按量付费肯定是不够的。

Heroku只有欧洲和美国的部署点，但也是git集成，非常方便，现在我已经转移到了heroku，主要我有github学生包，可以白嫖heroku

**heroku注册要关闭adguard，需要绑定国外银行卡，绑卡时建议关闭梯子**。

## Paas部署注意点

### 端口号

注意一下项目的端口号设置,最好设置在环境变量中,然后在项目中通过 `os.Getenv("xxx")`来获取端口号.

支持go部署的paas中，环境变量 `PORT` 一般都默认8080，且为全局的。无需写入环境变量，直接调用就好了。

- 对于tgbot：官方示例中使用的8443端口，在部署到paas平台时8443端口需要确认是否开放。

### tgbot证书问题

heroku和zeabur部署项目自带证书,做完域名映射可以直接https访问.
*所以在设置webhook进行和tg服务器通讯的时候不需要手动加载 `cert.pem` 和 `key.pem`*

在部署tg的bot时,可以修改tgbot官方对go语言搭建bot示例中的:

```go
  ...

  log.Printf("Authorized on account %s", bot.Self.UserName)
  wh, _ := tgbotapi.NewWebhook(TG_WEBHOOK_URL + bot.Token)

  ...

  go http.ListenAndServe(":"+port, nil)
```

直接使用 `NewWebhook`和 `ListenAndServe`函数即可.

### 环境变量

实例定义四个环境变量

```env
TG_WEBHOOK_URL=https://yousite.com/tgbot
BOT_TOKEN=your token
chatGPTAPIKey=sk-your key
ChatGPTURL_proxy = "https://your-proxy-site/v1"
```

tips: 免费的gpt api的代理 -> ChatGPTURL_proxy="[https://one-api.bltcy.top/v1](https://one-api.bltcy.top/v1)"
