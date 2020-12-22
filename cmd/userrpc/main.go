package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"pkg.zpf.com/golang/kit-scaffold/internal/app/userrpc/dao"
	"pkg.zpf.com/golang/kit-scaffold/internal/app/userrpc/endpoint"
	"pkg.zpf.com/golang/kit-scaffold/internal/app/userrpc/service"
	"pkg.zpf.com/golang/kit-scaffold/internal/app/userrpc/transport"
	"pkg.zpf.com/golang/kit-scaffold/internal/pkg/config"
	"pkg.zpf.com/golang/kit-scaffold/internal/pkg/databases"
	zaplog "pkg.zpf.com/golang/kit-scaffold/internal/pkg/log"
	"pkg.zpf.com/golang/kit-scaffold/internal/pkg/registers"
	pb "pkg.zpf.com/golang/kit-scaffold/protos/user"

	kitlog "github.com/go-kit/kit/log"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "userrpc.yaml", "set config file which will loading.")
var quitChan = make(chan error, 1)

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
	registrar, err := registers.InitRegister(
		ctx, kitlog.NewNopLogger(), config.Conf.EtcdConfig,
	)

	userDao := dao.NewUserDaoImpl(logger)
	userService := service.NewUserRPCServiceImpl(userDao, logger)
	userEndpoints := endpoint.UserRPCEndpoints{
		FindByIDEndpoint:    endpoint.MakeFindByIDEndpoint(userService),
		FindByEmailEndpoint: endpoint.MakeFindByEmailEndpoint(userService),
	}
	//将服务地址注册到etcd中
	go func() {
		registrar.Register()
		// 使用 transport 构造 UserService
		handler := transport.NewUserServer(ctx, userEndpoints)
		// 监听端口，建立 gRPC 网络服务器，注册 RPC 服务
		ls, err := net.Listen("tcp", config.Conf.EtcdConfig.GrpcAddr)
		if err != nil {
			logger.Error("listen tcp err", zap.Error(err))
			quitChan <- err
			return
		}
		gRPCServer := grpc.NewServer()
		pb.RegisterUserServer(gRPCServer, handler)
		err = gRPCServer.Serve(ls)
		if err != nil {
			logger.Error("gRPCServer Serve err", zap.Error(err))
			quitChan <- err
			return
		}
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
		quitChan <- fmt.Errorf("%s", <-c)
	}()
	err = <-quitChan
	registrar.Deregister()
	logger.Error("quit err", zap.Error(err))
}
