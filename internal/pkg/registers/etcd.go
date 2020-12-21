package registers

import (
	"blogger-kit/internal/pkg/config"
	"context"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"go.uber.org/zap"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/etcdv3"
)

// InitRegister etcd初始化注册服务
func InitRegister(ctx context.Context, logger log.Logger, cfg *config.EtcdConfig) (registrar sd.Registrar, err error) {
	client := connectEtcd(ctx, cfg)
	// 创建注册器
	registrar = etcdv3.NewRegistrar(
		client, etcdv3.Service{
			Key:   fmt.Sprintf("%s/%s", cfg.SerName, cfg.GrpcAddr),
			Value: cfg.GrpcAddr,
		},
		logger,
	)

	return
}

// ClientInstance etcd客户端实例
func ClientInstance(ctx context.Context, logger log.Logger, cfg *config.EtcdConfig) (*etcdv3.Instancer, error) {
	client := connectEtcd(ctx, cfg)
	instance, err := etcdv3.NewInstancer(client, cfg.SerName, logger)
	if err != nil {
		return instance, err
	}
	return instance, nil
}

// connectEtcd 连接etcd
func connectEtcd(ctx context.Context, cfg *config.EtcdConfig) etcdv3.Client {
	//etcd的连接参数
	options := etcdv3.ClientOptions{
		DialTimeout:   time.Second * cfg.TTL,
		DialKeepAlive: time.Second * cfg.TTL,
	}
	//创建etcd连接
	etcdServers := strings.Split(cfg.EtcdAddr, ",")
	client, err := etcdv3.NewClient(ctx, etcdServers, options)
	if err != nil {
		panic(err)
	}
	return client
}

// GetEntrie 获取etcd的grpc服务地址
func GetEntrie(ctx context.Context, logger *zap.Logger, cfg *config.EtcdConfig) (string, error) {
	client := connectEtcd(ctx, cfg)
	entries, err := client.GetEntries(cfg.SerName)
	if err != nil {
		logger.Error("获取etcd中grpc地址失败", zap.Error(err))
		return "", err
	}
	return entries[rand.Intn(len(entries))], nil
}
