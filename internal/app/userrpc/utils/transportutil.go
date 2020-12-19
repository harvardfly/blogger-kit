package utils

import (
	userpb "blogger-kit/protos/user"
	"context"
)

func DecodeFindByEmailRequest(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*userpb.FindByEmailRequest)
	return &userpb.FindByEmailRequest{
		Email: req.Email,
	}, nil
}

func EncodeFindByEmailResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*userpb.UserResponse)
	return &userpb.UserResponse{
		Id:       resp.Id,
		Username: resp.Username,
		Email:    resp.Email,
	}, nil
}

func DecodeFindByIDRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*userpb.FindByIDRequest)
	return &userpb.FindByIDRequest{
		Id: req.Id,
	}, nil
}

func EncodeFindByIDResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*userpb.UserResponse)
	return &userpb.UserResponse{
		Id:       resp.Id,
		Username: resp.Username,
		Email:    resp.Email,
	}, nil
}
