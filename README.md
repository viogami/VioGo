# VioGo

![Go Reference](https://pkg.go.dev/badge/github.com/go-telegram-bot-api/telegram-bot-api/v5.svg)
![viogo文档](https://img.shields.io/badge/version-1.0-violet)

<p align="center"><img src="./docs/logo.png" width="120" height="120" alt="logo"></p>

<p align="center"><a href="https://viogami.github.io/VioGo/">中文文档</a> • <a href="#快速开始">快速开始</a> </p>

golang编写的机器人业务后端，bot取名为 **Vio** ，qq机器人的实现基于gocq提供的api，监听gocq的上报事件调用转发外部讯息并返回。

可以完成基于chatgpt聊天任务，以及一些简单的指令响应。

**目前正在重构，并且开发更多后台命令！除了聊天,VioGo的目标是可以在后台分析聊天数据的奇妙bot！**

**如果你有想法和意见请提issue！这对我和bot都很重要！**
**欢迎qq加群讨论：340961300**

实现平台以及实现的外部功能:

- [X] QQ bot
  - [X] 随机涩图
  - [X] 猎杀对决枪声语音
  - [X] 发送已撤回的消息
  - [X] 禁言抽奖

- 天然支持http请求，使用go原生net包，创建了 `/post`请求的路由，可以解析post内容转发调用chatgpt。
- 配合[go-cqhttp](https://github.com/Mrs4s/go-cqhttp)使用，用于实现qqbot,基于上报事件，调用gocq提供的api进行回复。注意参考[gocq官方文档](https://docs.go-cqhttp.org/reference/#websocket)。可以调用api发送图片，随机图片api[参考](https://api.lolicon.app/setu/v2)

## 快速开始

在安装本bot之前你需要部署两个前置服务：gocq服务以及qsign签名服务器

- gocq可以通过[官方途径](https://github.com/Mrs4s/go-cqhttp)部署。gocq提供了链接指定bot的qq号，本项目只提供了bot的消息处理逻辑。
- qsign服务器可以通过docker拉取，建议拉取新更新的，请使用txlib 9.x.x版本和手表协议，不然容易报错低版本。
- 更多细节可以参考我的[个人博客](http://viogami.tech/index.php/blog/144/)

### 部署到paas平台

可fork本仓库，用git连接到指定paas平台部署，注意配置环境变量即可。

环境变量参考 `config/env.go`文件

### 部署到本地

clone仓库源码
**需要本地go环境**
打包成可执行文件

以linux为例：

```bash
set goos=linux
go build main.go
```

之后配置环境变量挂载指定端口即可启动
