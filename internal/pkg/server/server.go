package server

import (
	"blogger-kit/internal/pkg/config"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

// InitServer 初始化http server
func InitServer(cfg *config.ServerConfig, r http.Handler) (err error) {
	errChan := make(chan error)
	go func() {
		errChan <- http.ListenAndServe(":"+strconv.Itoa(cfg.Port), r)
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
