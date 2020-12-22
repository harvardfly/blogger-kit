package transport

import (
	"context"

	"pkg.zpf.com/golang/kit-scaffold/internal/app/userrpc/endpoint"
	"pkg.zpf.com/golang/kit-scaffold/internal/app/userrpc/utils"
	userpb "pkg.zpf.com/golang/kit-scaffold/protos/user"

	kitrpc "github.com/go-kit/kit/transport/grpc"
)

// grpcServer 定义grpc server
type grpcServer struct {
	findByEmail kitrpc.Handler
	findByID    kitrpc.Handler
}

// FindByEmail 实现 userpb.UserServer
func (g grpcServer) FindByEmail(ctx context.Context, r *userpb.FindByEmailRequest) (*userpb.UserResponse, error) {
	_, resp, err := g.findByEmail.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*userpb.UserResponse), nil
}

// FindById 实现 userpb.UserServer
func (g grpcServer) FindById(ctx context.Context, r *userpb.FindByIDRequest) (*userpb.UserResponse, error) {
	_, resp, err := g.findByID.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*userpb.UserResponse), nil
}

// NewUserServer 初始化UserServer
func NewUserServer(ctx context.Context, endpoints endpoint.UserRPCEndpoints) userpb.UserServer {
	return &grpcServer{
		findByEmail: kitrpc.NewServer(
			endpoints.FindByEmailEndpoint,
			utils.DecodeFindByEmailRequest,
			utils.EncodeFindByEmailResponse,
		),
		findByID: kitrpc.NewServer(
			endpoints.FindByIDEndpoint,
			utils.DecodeFindByIDRequest,
			utils.EncodeFindByIDResponse,
		),
	}
}
