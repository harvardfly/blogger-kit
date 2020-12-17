package component

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/etcdv3"
)

func Register(ctx context.Context, etcdServer, svcHost, svcPort string, logger log.Logger) (registar sd.Registrar) {
	//etcd的连接参数
	options := etcdv3.ClientOptions{
		DialTimeout:   time.Second * 3,
		DialKeepAlive: time.Second * 3,
	}
	//创建etcd连接
	client, err := etcdv3.NewClient(ctx, []string{etcdServer}, options)
	if err != nil {
		panic(err)
	}

	// 创建注册器
	registrar := etcdv3.NewRegistrar(client, etcdv3.Service{
		Key:   fmt.Sprintf("%s:%s", svcHost, svcPort),
		Value: svcPort,
	}, logger)

	// 注册器启动注册
	registrar.Register()

	return
}
