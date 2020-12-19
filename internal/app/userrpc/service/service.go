package service

import (
	"blogger-kit/internal/app/userrpc/dao"
	userpb "blogger-kit/protos/user"
	"context"
	"errors"

	"go.uber.org/zap"
)

var (
	// ErrNotFound 定义错误
	ErrNotFound = errors.New("用户不存在")
)

// UserService 用户服务
type UserRPCService interface {
	FindByEmail(ctx context.Context, req *userpb.FindByEmailRequest) (*userpb.UserResponse, error)
	FindByID(ctx context.Context, req *userpb.FindByIDRequest) (*userpb.UserResponse, error)
}

// UserRPCServiceImpl 初始默认的
type UserRPCServiceImpl struct {
	userDao dao.UserDao
	logger  *zap.Logger
}

// NewUserRPCServiceImpl 初始化
func NewUserRPCServiceImpl(userDao dao.UserDao, logger *zap.Logger) UserRPCService {
	return &UserRPCServiceImpl{
		userDao: userDao,
		logger:  logger.With(zap.String("type", "NewUserRPCServiceImpl")),
	}
}

// FindByEmail 实现rpc服务通过Token获取用户信息
func (s *UserRPCServiceImpl) FindByEmail(ctx context.Context, req *userpb.FindByEmailRequest) (*userpb.UserResponse, error) {
	member, err := s.userDao.FindByEmail(req.Email)
	if err != nil {
		return &userpb.UserResponse{}, ErrNotFound
	}
	return &userpb.UserResponse{
		Id:       int32(member.ID),
		Username: member.Username,
		Email:    member.Email,
	}, nil
}

// FindByID 调用rpc服务通过ID获取用户信息
func (s *UserRPCServiceImpl) FindByID(ctx context.Context, req *userpb.FindByIDRequest) (*userpb.UserResponse, error) {
	member, err := s.userDao.FindByID(req.Id)
	if err != nil {
		return &userpb.UserResponse{}, ErrNotFound
	}
	return &userpb.UserResponse{
		Id:       int32(member.ID),
		Username: member.Username,
		Email:    member.Email,
	}, nil
}
