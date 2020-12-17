# blogger-kit
基于go-kit的微服务架构脚手架

blogger-kit #项目根路径
├── deploy #Dockerfile等
├── docs #整体文档
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