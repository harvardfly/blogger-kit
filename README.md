# blogger-kit
基于go-kit的微服务架构脚手架,集成zap日志库、mysql、redis、grpc、etcd等常用组件，可作为项目脚手架开箱即用；项目部署采用k8s和docker swarm

## 项目结构
```markdown
kit-scaffold #项目根路径
├── cmd #服务启动main文件
├── conf #配置文件
├── deploy #Dockerfile等
├── docs #整体文档
├── protos #rpc protobuf文件
├── internal #项目模块
│   ├── app # 具体的服务应用
│   │   ├── user #用户服务示例
│   │   │   ├── config #配置文件
│   │   │   ├── dao #数据层
│   │   │   ├── endpoint #逻辑层
│   │   │   ├── server #服务层
│   │   │   ├── test #单元测试示例
│   │   │   ├── utils #应用内部工具方法
│   │   │   └── transport #路由 提供http、rpc
├── ├── pkg #公共工具
│   │   ├── component #服务注册组件
│   │   └── database #数据库类   
│   │   ├── log #日志模块
│   │   ├── models #数据表
│   │   ├── redis #redis模块
│   │   ├── requests #请求struct
│   │   ├── response #响应struct
|   |   ├── utils #项目公共工具方法
└── vendor #go mod require
```

## 集成组件
```markdown
目前已集成：
gin作为路由
config配置文件
MySQL数据库
redis缓存
http/grpc传输
zap日志库
jwt认证中间件
etcd服务注册发现
docker swam
k8s
```

## 系统环境要求
```$xslt
golang >= 1.13
```

## 项目部署
### docker swam方式：
```markdown
1. make build 生成二进制文件
2. make docker 打包镜像并推送到docker hub
3. 执行docker-compose.yaml 启动服务
```
### k8s方式：
```markdown
1、2步同docker swam
3. 执行batch_deploy.sh启动pod
执行batch_undeploy.sh可关闭pod
```

## golint 代码规范检查
```$xslt
1. cd blogger-kit
2. go list ./... | grep -v /vendor/ | xargs -L1 golint -set_exit_status
```