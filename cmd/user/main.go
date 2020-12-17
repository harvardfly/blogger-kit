package main

import (
	"blogger-kit/internal/app/user/dao"
	"blogger-kit/internal/app/user/endpoint"
	"blogger-kit/internal/app/user/service"
	"blogger-kit/internal/app/user/transport"
	"blogger-kit/internal/pkg/config"
	"blogger-kit/internal/pkg/databases"
	zaplog "blogger-kit/internal/pkg/log"
	"blogger-kit/internal/pkg/redis"
	"blogger-kit/internal/pkg/server"
	"context"
	"flag"
	"log"
)

var configFile = flag.String("f", "user.yaml", "set config file which will loading.")

func main() {
	flag.Parse()
	err := config.Init(*configFile)
	if err != nil {
		panic(err)
	}

	// 初始化日志模块
	logger, err := zaplog.InitLog(config.Conf.LogConfig)
	if err != nil {
		log.Println("加载日志配置失败")
	}

	// 初始化数据库模块
	err = databases.InitMysql(config.Conf.MySQLConfig)
	if err != nil {
		log.Println("加载数据库配置失败")
	}

	// 初始化redis模块
	err = redis.InitRedis(config.Conf.RedisConfig)
	if err != nil {
		log.Println("加载redis配置失败")
	}
	ctx := context.Background()
	userDAO := dao.NewUserDAOImpl(logger)
	userService := service.NewUserServiceImpl(userDAO, logger)
	userEndpoints := &endpoint.UserEndpoints{
		RegisterEndpoint: endpoint.MakeRegisterEndpoint(userService),
		LoginEndpoint:    endpoint.MakeLoginEndpoint(userService),
	}

	// 初始化Server
	r := transport.NewHttpHandler(ctx, userEndpoints, logger)
	server.InitServer(config.Conf.ServerConfig, r)
}
