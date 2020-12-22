package utils

import (
	"context"

	userpb "pkg.zpf.com/golang/kit-scaffold/protos/user"
)

// DecodeFindByEmailRequest  FindByEmail参数验证
func DecodeFindByEmailRequest(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*userpb.FindByEmailRequest)
	return &userpb.FindByEmailRequest{
		Email: req.Email,
	}, nil
}

// EncodeFindByEmailResponse  FindByEmail返回值
func EncodeFindByEmailResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*userpb.UserResponse)
	return &userpb.UserResponse{
		Id:       resp.Id,
		Username: resp.Username,
		Email:    resp.Email,
	}, nil
}

// DecodeFindByIDRequest FindByID 参数验证
func DecodeFindByIDRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*userpb.FindByIDRequest)
	return &userpb.FindByIDRequest{
		Id: req.Id,
	}, nil
}

// EncodeFindByIDResponse FindByID 返回值验证
func EncodeFindByIDResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*userpb.UserResponse)
	return &userpb.UserResponse{
		Id:       resp.Id,
		Username: resp.Username,
		Email:    resp.Email,
	}, nil
}
