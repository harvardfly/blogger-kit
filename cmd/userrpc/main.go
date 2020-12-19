package main

import (
	"blogger-kit/internal/app/userrpc/dao"
	"blogger-kit/internal/app/userrpc/endpoint"
	"blogger-kit/internal/app/userrpc/service"
	"blogger-kit/internal/app/userrpc/transport"
	"blogger-kit/internal/pkg/config"
	"blogger-kit/internal/pkg/databases"
	zaplog "blogger-kit/internal/pkg/log"
	"context"
	"flag"
	"log"
	"net"

	pb "blogger-kit/protos/user"

	"google.golang.org/grpc"
)

var configFile = flag.String("f", "userrpc.yaml", "set config file which will loading.")

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

	ctx := context.Background()
	userDao := dao.NewUserDaoImpl(logger)
	userService := service.NewUserRPCServiceImpl(userDao, logger)
	userEndpoints := endpoint.UserRPCEndpoints{
		FindByIDEndpoint:    endpoint.MakeFindByIDEndpoint(userService),
		FindByEmailEndpoint: endpoint.MakeFindByEmailEndpoint(userService),
	}

	// 使用 transport 构造 UserService
	handler := transport.NewUserServer(ctx, userEndpoints)
	// 监听端口，建立 gRPC 网络服务器，注册 RPC 服务
	ls, _ := net.Listen("tcp", "127.0.0.1:8080")
	gRPCServer := grpc.NewServer()
	pb.RegisterUserServer(gRPCServer, handler)
	gRPCServer.Serve(ls)
}
