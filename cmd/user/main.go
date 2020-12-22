package main

import (
	"context"
	"flag"
	"log"

	"pkg.zpf.com/golang/kit-scaffold/internal/app/user/dao"
	"pkg.zpf.com/golang/kit-scaffold/internal/app/user/endpoint"
	"pkg.zpf.com/golang/kit-scaffold/internal/app/user/service"
	"pkg.zpf.com/golang/kit-scaffold/internal/app/user/transport"
	"pkg.zpf.com/golang/kit-scaffold/internal/pkg/config"
	"pkg.zpf.com/golang/kit-scaffold/internal/pkg/databases"
	zaplog "pkg.zpf.com/golang/kit-scaffold/internal/pkg/log"
	"pkg.zpf.com/golang/kit-scaffold/internal/pkg/redis"
	"pkg.zpf.com/golang/kit-scaffold/internal/pkg/registers"
	"pkg.zpf.com/golang/kit-scaffold/internal/pkg/server"
	pb "pkg.zpf.com/golang/kit-scaffold/protos/user"

	"go.uber.org/zap"

	"google.golang.org/grpc"
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
	// 初始化rpc客户端
	//instance, err := registers.ClientInstance(
	//	ctx,
	//	kitlog.NewNopLogger(),
	//	config.Conf.EtcdConfig,
	//)
	entrie, err := registers.GetEntrie(
		ctx,
		logger,
		config.Conf.EtcdConfig,
	)
	conn, err := grpc.Dial(entrie, grpc.WithInsecure())
	if err != nil {
		logger.Error("连接user rpc 错误", zap.Error(err))
		panic("grpc connect error")
	}
	defer conn.Close()
	userClient := pb.NewUserClient(conn)

	userDao := dao.NewUserDaoImpl(logger)
	userService := service.NewUserServiceImpl(userDao, logger, userClient)
	userEndpoints := &endpoint.UserEndpoints{
		RegisterEndpoint:    endpoint.MakeRegisterEndpoint(userService),
		LoginEndpoint:       endpoint.MakeLoginEndpoint(userService),
		FindByIDEndpoint:    endpoint.MakeFindByIDEndpoint(userService),
		FindByEmailEndpoint: endpoint.MakeFindByEmailEndpoint(userService),
	}

	// 初始化Server
	ctx = context.WithValue(ctx, "ginMod", config.Conf.Mode)
	r := transport.NewHTTPHandler(ctx, userEndpoints, logger)
	err = server.InitServer(config.Conf.ServerConfig, r)
	if err != nil {
		logger.Error("服务启动错误", zap.Error(err))
	}
}
