# GoBot 

![Go Reference](https://pkg.go.dev/badge/github.com/go-telegram-bot-api/telegram-bot-api/v5.svg)

使用go搭建的bot，取名为 **Vio** ，旨在提供一个接口，用来完成chatgpt聊天等任务，部署在一个服务器可以多个平台共同调用。

实现平台:
 - [x] Telegram bot
 - [ ] QQ bot

## 写在前面/preface
有部署聊天机器人的想法，但是我使用的国内服务器，而且服务器性能也堪忧，于是决定不要laas--即用云服务器部署了，找个国外的Paas平台，把写的后端送上去就好了，而且一般也都有免费计划，够用了。

## 对比表格/compare popular Paas
| 服务提供商  | Fly.io          | Railway        | Render         | Glitch         | Adaptable      | **Zeabur**      |
|-------------|-----------------|----------------|----------------|----------------|----------------|----------------|
| 长时间不活动关闭 | 否             | 否              | 15 分钟         | 5 分钟          | 是*            | 否             |
| 需要信用卡    | 是             | 是             | 否             | 否             | 否             | 否             |
| 免费计划      | 免费创建三个最低配应用| 试用$5额度| 750 小时       | 1000 小时      | 无*            | $5/mo        |
| 内存          | 256MB          | 512MB          | 512MB          | 512MB          | 256MB          | 512MB          |
| 磁盘空间      | 3GB            | 1GB            | -              | 200MB*         | 1GB            | 1GB            |
| 可写磁盘      | 是             | -              | 否             | 是             | 是*            | 是             |
| 网络带宽      | 160GB          | $0.10/GB       | 100GB          | 4000 次请求/时  | 100GB          | -              |
| 可用 Dockerfile | 是             | 是             | 是             | 否             | 否*            | 是             |
| GitHub 集成  | 是            | 是             | 是             | 是             | 是             | 是             |

对于
 - Cloudflare Workers 
 - AirCode（国内团队做的） 
 - Vercel
它们主要侧重于
>Fullstack Javascript Apps - Deploy and Host in Seconds

对非 Nodejs 的后端参考意义不大。
主要思路都是 Edge Network + Serverless Functions（函数代码在轻量级的 V8 沙盒中执行）

## 对于我/for me
只使用过cloudflare的workers部署服务，但是只能使用js，不熟悉还是挺难搞了，但是用js的可能比较舒服。
应该是完全免费的，只要绑了自己的域名在CF上，CF也提供子域名

目前感觉Zeabur挺不错，~~主要看重免费计划~~。国内社区，discord上回复也很即使，一键部署挺快的，github集成。

只部署一个机器人接口就好了.无论什么聊天平台，通讯功能的实现基本都是互通的。

本后端最终希望实现只对外暴露一个API,实现机器人通讯的应答模式,对不同平台创建不同的新服务,调用接口皆可进行通讯服务.

## zeabur部署注意点
### 端口号
zeabur上项目部署非常快,甚至不用写dockfile,而且对go项目有完整的支持,算是符合他们的口号:
> Deploying your service with one click

但是注意一下项目的端口号设置,最好设置在环境变量中,然后在项目中通过`os.Getenv("xxx")`来获取端口号.

zeabur的go项目中，环境变量`PORT`是默认8080，且为全局的。也可以不设置，直接调用就好了。

### 证书问题
zeabur部署项目自带证书,做完域名映射可以直接https访问.
所以在设置webhook进行和tg服务器通讯的时候不需要手动加载`cert.pem`和`key.pem`

在部署tg的bot时,可以修改tgbot官方对go语言搭建bot示例中的:
``` go
  ...

  log.Printf("Authorized on account %s", bot.Self.UserName)
	wh, _ := tgbotapi.NewWebhook(TG_WEBHOOK_URL + bot.Token)

  ...

  go http.ListenAndServe(":"+port, nil)
```
直接使用`NewWebhook`和`ListenAndServe`函数即可.

--------------
--------------
表格和paas平台对比参考：[免费的 PaaS 平台汇总][1]

  [1]: https://liduos.com/Summary-of-free-PaaS-platforms.html

