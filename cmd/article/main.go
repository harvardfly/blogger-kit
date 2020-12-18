package main

import (
	"blogger-kit/internal/app/article/dao"
	"blogger-kit/internal/app/article/endpoint"
	"blogger-kit/internal/app/article/service"
	"blogger-kit/internal/app/article/transport"
	"blogger-kit/internal/pkg/config"
	"blogger-kit/internal/pkg/databases"
	zaplog "blogger-kit/internal/pkg/log"
	"blogger-kit/internal/pkg/redis"
	"blogger-kit/internal/pkg/server"
	"context"
	"flag"
	"log"
)

var configFile = flag.String("f", "article.yaml", "set config file which will loading.")

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
	articleDao := dao.NewArticleDaoImpl(logger)
	articleService := service.NewArticleServiceImpl(articleDao, logger)
	articleEndpoints := &endpoint.ArticleEndpoints{
		ArticleEndpoint:    endpoint.MakeArticleEndpoint(articleService),
		GetArticleEndpoint: endpoint.MakeGetArticleEndpoint(articleService),
	}

	// 初始化Server
	r := transport.NewHttpHandler(ctx, articleEndpoints, logger)
	server.InitServer(config.Conf.ServerConfig, r)
}
