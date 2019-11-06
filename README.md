## 特性

- 遵循 RESTful API 设计规范
- 包含3个微服务(api, user, auth)
- 微服务间使用grpc通信
- 依赖注入(基于[wire](https://github.com/google/wire))
- JWT 认证(基于黑名单的认证模式，存储支持：file(buntdb内存数据库))

## 项目结构概览

```shell
.
├── build      微服务镜像构建Dockerfile
│   ├── api
│   ├── auth
│   └── user
├── build.sh
├── cmd        微服务入口
│   ├── api
│   ├── auth
│   └── user
├── configs
│   └── api.toml
├── docker-compose.yml      本地docker-compose运行
├── internal                内部应用
│   └── app
├── Makefile
├── pkg                     公共库
│   ├── app
│   ├── errors
│   ├── log
│   └── util
├── proto                   微服务协议文件
│   ├── auth
│   ├── proto.go
│   └── user
├── README.md
└── postman：postman api调试request集合
```

## 本地测试

### 环境需求:

* golang 1.13
* docker, docker-compose

### 本地运行

* `make build`
* `make docker`

### 本地测试

使用提供的postman request集合进行api测试

### 清理环境

`make clean-image`

## Todo List

- [ ] 添加邮件重制密码api
- [ ] 微服务统一配置
- [ ] 微服务统一错误码
- [ ] 微服务统一日志
- [ ] 服务注册由consul替换成k8s
- [ ] 删除冗余代码及代码层次进一步分离