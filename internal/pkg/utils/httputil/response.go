package httputil

import (
	"blogger-kit/internal/pkg/baseerror"
	"blogger-kit/internal/pkg/utils/middlewareutil"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"go.uber.org/zap"
)

func Error(err error, msg string) interface{} {
	return struct {
		Err string `json:"err"`
		Msg string `json:"msg"`
	}{
		Err: err.Error(),
		Msg: msg,
	}
}

func Success(data interface{}) interface{} {
	return struct {
		Data interface{} `json:"data"`
	}{
		Data: data,
	}
}

func ReturnJson(code int64, message string, data interface{}) interface{} {
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

func EncodeJSONResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

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
