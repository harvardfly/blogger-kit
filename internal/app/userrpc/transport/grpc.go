package transport

import (
	"blogger-kit/internal/app/userrpc/endpoint"
	"blogger-kit/internal/app/userrpc/utils"
	userpb "blogger-kit/protos/user"
	"context"

	kitrpc "github.com/go-kit/kit/transport/grpc"
)

type grpcServer struct {
	findByEmail kitrpc.Handler
	findByID    kitrpc.Handler
}

// FindByEmail
func (g grpcServer) FindByEmail(ctx context.Context, r *userpb.FindByEmailRequest) (*userpb.UserResponse, error) {
	_, resp, err := g.findByEmail.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*userpb.UserResponse), nil
}

func (g grpcServer) FindById(ctx context.Context, r *userpb.FindByIDRequest) (*userpb.UserResponse, error) {
	_, resp, err := g.findByID.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*userpb.UserResponse), nil
}

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
