package endpoint

import (
	"context"

	"pkg.zpf.com/golang/kit-scaffold/internal/app/user/service"
	"pkg.zpf.com/golang/kit-scaffold/internal/pkg/requests"
	"pkg.zpf.com/golang/kit-scaffold/internal/pkg/responses"

	"github.com/go-kit/kit/endpoint"
)

// UserEndpoints 用户模块Endpoints
type UserEndpoints struct {
	RegisterEndpoint    endpoint.Endpoint
	LoginEndpoint       endpoint.Endpoint
	FindByIDEndpoint    endpoint.Endpoint
	FindByEmailEndpoint endpoint.Endpoint
}

// MakeLoginEndpoint 登录Endpoint
func MakeLoginEndpoint(userService service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (res interface{}, err error) {
		logReq := request.(*requests.LoginRequest)
		loginInfo, err := userService.Login(ctx, logReq.Email, logReq.Password)
		return loginInfo, err
	}
}

// MakeRegisterEndpoint 注册用户Endpoint
func MakeRegisterEndpoint(userService service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (res interface{}, err error) {
		regReq := request.(*requests.RegisterRequest)
		userInfo, err := userService.Register(ctx, &responses.RegisterUser{
			Username: regReq.Username,
			Password: regReq.Password,
			Email:    regReq.Email,
		})
		if err != nil {
			return nil, err
		}
		return &responses.RegisterResponse{UserInfo: userInfo}, nil
	}
}

// MakeFindByIDEndpoint ID获取用户信息Endpoint
func MakeFindByIDEndpoint(userService service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (res interface{}, err error) {
		req := request.(*requests.FindByIDRequest)
		loginInfo, err := userService.FindByID(ctx, req)
		return loginInfo, err
	}
}

// MakeFindByEmailEndpoint 邮箱获取用户信息Endpoint
func MakeFindByEmailEndpoint(userService service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (res interface{}, err error) {
		req := request.(*requests.FindByEmailRequest)
		userInfo, err := userService.FindByEmail(ctx, req)
		if err != nil {
			return nil, err
		}
		return userInfo, nil
	}
}
