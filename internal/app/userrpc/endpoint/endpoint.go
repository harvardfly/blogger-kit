package endpoint

import (
	"context"

	"pkg.zpf.com/golang/blogger-kit/internal/app/userrpc/service"
	userpb "pkg.zpf.com/golang/blogger-kit/protos/user"

	"github.com/go-kit/kit/endpoint"
)

// UserRPCEndpoints 用户RPC模块Endpoints
type UserRPCEndpoints struct {
	FindByEmailEndpoint endpoint.Endpoint
	FindByIDEndpoint    endpoint.Endpoint
}

// MakeFindByEmailEndpoint 通过邮箱查询用户Endpoint
func MakeFindByEmailEndpoint(userRPCService service.UserRPCService) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (res interface{}, err error) {
		userReq := req.(*userpb.FindByEmailRequest)
		userInfo, err := userRPCService.FindByEmail(ctx, userReq)
		return userInfo, err
	}
}

// MakeFindByIDEndpoint 通过ID查询用户Endpoint
func MakeFindByIDEndpoint(userRPCService service.UserRPCService) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (res interface{}, err error) {
		userReq := req.(*userpb.FindByIDRequest)
		userInfo, err := userRPCService.FindByID(ctx, userReq)
		if err != nil {
			return nil, err
		}
		return userInfo, nil
	}
}
