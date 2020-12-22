package main

import (
	"context"
	"flag"
	"log"

	"pkg.zpf.com/golang/kit-scaffold/internal/app/article/dao"
	"pkg.zpf.com/golang/kit-scaffold/internal/app/article/endpoint"
	"pkg.zpf.com/golang/kit-scaffold/internal/app/article/service"
	"pkg.zpf.com/golang/kit-scaffold/internal/app/article/transport"
	"pkg.zpf.com/golang/kit-scaffold/internal/pkg/config"
	"pkg.zpf.com/golang/kit-scaffold/internal/pkg/databases"
	zaplog "pkg.zpf.com/golang/kit-scaffold/internal/pkg/log"
	"pkg.zpf.com/golang/kit-scaffold/internal/pkg/redis"
	"pkg.zpf.com/golang/kit-scaffold/internal/pkg/server"
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
	ctx = context.WithValue(ctx, "ginMod", config.Conf.Mode)
	r := transport.NewHTTPHandler(ctx, articleEndpoints, logger)
	server.InitServer(config.Conf.ServerConfig, r)
}
