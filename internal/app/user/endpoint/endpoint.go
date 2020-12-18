package endpoint

import (
	"blogger-kit/internal/app/user/service"
	"blogger-kit/internal/pkg/requests"
	"blogger-kit/internal/pkg/responses"
	"context"

	"github.com/go-kit/kit/endpoint"
)

// UserEndpoints 用户模块Endpoints
type UserEndpoints struct {
	RegisterEndpoint endpoint.Endpoint
	LoginEndpoint    endpoint.Endpoint
}

// MakeLoginEndpoint 登录Endpoint
func MakeLoginEndpoint(userService service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (res interface{}, err error) {
		logReq := req.(*requests.LoginRequest)
		loginInfo, err := userService.Login(ctx, logReq.Email, logReq.Password)
		return loginInfo, err
	}
}

// MakeRegisterEndpoint 注册用户Endpoint
func MakeRegisterEndpoint(userService service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (res interface{}, err error) {
		regReq := req.(*requests.RegisterRequest)
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
