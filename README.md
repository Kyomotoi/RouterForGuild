# RouterForGuild

一个高效的，为`Tencent Guild（腾讯官方频道机器人）`而生的Router。目前正不断完善中...</br>
由于Owner学业问题，完善该项目可能需要很长一段时间（2022.6之前）

## 声明

**一切开发旨在学习，请勿用于非法用途**

## 特性

- 对所发/收消息进行GFW（非法敏感词）筛查
- 内置黑名单，可在处理消息前隔绝
- 内置管理功能，可对bot主体（插件、限制、黑名单等）直接控制
- 插件通过`正向WebSocket`连接，以实现`热插拔、“分布式”`的目的
- 插件标准可选`OneBot 群聊`协议，直接兼容基于`OneBot v12<=`此前的bot

>如果你不知道什么是`OneBot标准`，请查阅[此处](https://github.com/botuniverse/onebot)

## 架构

![架构](https://cdn.jsdelivr.net/gh/Kyomotoi/CDN@master/noting/structure.png)

## 鸣谢

[BotUniverse](https://github.com/botuniverse): [OneBot](https://github.com/botuniverse/onebot) <br>
[wdvxdr1123](https://github.com/wdvxdr1123): [ZeroBot](https://github.com/wdvxdr1123/ZeroBot)

## 协议
本项目使用[AGPLv3](LICENSE)
