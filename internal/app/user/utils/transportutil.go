package utils

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"pkg.zpf.com/golang/blogger-kit/internal/pkg/baseerror"

	"pkg.zpf.com/golang/blogger-kit/internal/pkg/requests"
)

var (
	// ErrorBadRequest 参数错误
	ErrorBadRequest = errors.New("invalid request parameter")
)

// DecodeRegisterRequest 注册请求参数
func DecodeRegisterRequest(_ context.Context, r *http.Request) (interface{}, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("read body err, %v\n", err)
		return nil, err
	}
	var reg requests.RegisterRequest
	if err = json.Unmarshal(body, &reg); err != nil {
		log.Printf("Unmarshal err, %v\n", err)
		return nil, err
	}
	// 验证必填参数
	err = baseerror.ParamError(reg)
	if err != nil {
		return nil, err
	}
	// 自定义业务验证
	if reg.Username == "" || reg.Password == "" || reg.Email == "" {
		return nil, ErrorBadRequest
	}

	return &requests.RegisterRequest{
		Username: reg.Username,
		Email:    reg.Email,
		Password: reg.Password,
	}, nil
}

// DecodeLoginRequest  登录请求参数
func DecodeLoginRequest(_ context.Context, r *http.Request) (interface{}, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("read body err, %v\n", err)
		return nil, err
	}
	var reg requests.LoginRequest
	if err = json.Unmarshal(body, &reg); err != nil {
		log.Printf("Unmarshal err, %v\n", err)
		return nil, err
	}
	// 验证必填参数
	err = baseerror.ParamError(reg)
	if err != nil {
		return nil, err
	}
	if reg.Email == "" || reg.Password == "" {
		return nil, ErrorBadRequest
	}
	return &requests.LoginRequest{
		Email:    reg.Email,
		Password: reg.Password,
	}, nil
}

// DecodeFindByIDRequest 通过ID查找用户 请求参数
func DecodeFindByIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	id := r.URL.Query().Get("id")
	if id == "" {
		return nil, ErrorBadRequest
	}
	idInt, _ := strconv.Atoi(id)
	return &requests.FindByIDRequest{
		ID: idInt,
	}, nil
}
