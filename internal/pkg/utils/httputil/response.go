package httputil

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	ginZap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"

	"pkg.zpf.com/golang/blogger-kit/internal/pkg/baseerror"
	"pkg.zpf.com/golang/blogger-kit/internal/pkg/utils/middlewareutil"

	kithttp "github.com/go-kit/kit/transport/http"
	"go.uber.org/zap"
)

// InitControllers init controllers
type InitControllers func(r *gin.Engine)

// Error 定义错误返回
func Error(err error, msg string) interface{} {
	return struct {
		Err string `json:"err"`
		Msg string `json:"msg"`
	}{
		Err: err.Error(),
		Msg: msg,
	}
}

// Success 定义成功返回
func Success(data interface{}) interface{} {
	return struct {
		Data interface{} `json:"data"`
	}{
		Data: data,
	}
}

// ReturnJSON 返回json
func ReturnJSON(code int64, message string, data interface{}) interface{} {
	return struct {
		Code    int64       `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

// ServerOptions 初始化http server options
func ServerOptions(logger *zap.Logger) []kithttp.ServerOption {
	options := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(func(ctx context.Context, err error, w http.ResponseWriter) {
			logger.Warn(fmt.Sprint(ctx.Value("")), zap.Error(err))
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			json.NewEncoder(w).Encode(baseerror.ErrorWrapper{Error: err.Error()})
		}), //程序中的全部报错都会走这里面
		kithttp.ServerBefore(func(ctx context.Context, request *http.Request) context.Context {
			ctx = context.WithValue(ctx, middlewareutil.JWTConTextKey, request.Header.Get("Authorization"))
			logger.Debug("把请求中的token发到Context中", zap.Any("Token", request.Header.Get("Authorization")))
			return ctx
		}),
	}
	return options
}

// EncodeJSONResponse 返回json
func EncodeJSONResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

// EncodeError 错误处理
func EncodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func NewRouter(mode string, logger *zap.Logger) *gin.Engine {
	// 配置gin
	gin.SetMode(mode)
	r := gin.New()
	// panic之后自动恢复
	r.Use(gin.Recovery())
	// 日志格式化
	r.Use(ginZap.Ginzap(logger, time.RFC3339, true))
	// panic日志格式化
	r.Use(ginZap.RecoveryWithZap(logger, true))

	return r
}
