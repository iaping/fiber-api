# fiber-api

![docs](docs.png)

fiber-api是一个基于fiber的API基础开发框架。封装简单易懂，把更多时间放在业务开发上。

## 第三方
- [fiber](https://github.com/gofiber/fiber) web框架
- [bun](https://github.com/uptrace/bun) DB库
- [zerolog](https://github.com/rs/zerolog) 日志库
- [redis](https://github.com/redis/go-redis) redis缓存
- [swagger](https://github.com/gofiber/swagger) API文档
- [cron](https://github.com/robfig/cron) 定时任务

## 安装
```bash
git clone git@github.com:iaping/fiber-api.git
cd fiber-api
make
```
或自行go build

## API文档
请先安装[swag](https://github.com/swaggo/swag)
```bash
# swag init
make doc
```
访问：http://127.0.0.1:8080/_doc/index.html

## 说明

ctx.Ctx为整个系统的桥梁（全局资源的上下文）！初始常用的资源，如数据库、缓存的操作等。自己需要的资源自己注入并初始化即可，方便全局调用。

## 最后
如果对你有帮助，欢迎star或提交代码完善此项目。