package server

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/gin-gonic/gin"

	"pkg.zpf.com/golang/kit-scaffold/internal/pkg/config"
)

// InitServer 初始化http server
func InitServer(cfg *config.ServerConfig, r *gin.Engine) (err error) {
	errChan := make(chan error)
	go func() {
		errChan <- r.Run(fmt.Sprintf(":%s", strconv.Itoa(cfg.Port)))
	}()

	go func() {
		// 监控系统信号，等待 ctrl + c 系统信号通知服务关闭
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	err = <-errChan
	fmt.Println(err)
	return
}
