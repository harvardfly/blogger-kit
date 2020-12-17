package service

import (
	"blogger-kit/internal/app/user/dao"
	"blogger-kit/internal/pkg/models"
	"blogger-kit/internal/pkg/responses"
	"context"
	"errors"

	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

var (
	ErrUserExisted = errors.New("user is existed")
	ErrPassword    = errors.New("email and password are not match")
	ErrRegistering = errors.New("email is registering")
)

// UserService 用户服务
type UserService interface {
	// Login 登录接口
	Login(ctx context.Context, email, password string) (*responses.UserInfo, error)
	// Register 注册接口
	Register(ctx context.Context, vo *responses.RegisterUser) (*responses.UserInfo, error)
}

// UserServiceImpl 初始默认的
type UserServiceImpl struct {
	userDAO dao.UserDAO
	logger  *zap.Logger
}

// NewUserServiceImpl 初始化
func NewUserServiceImpl(userDAO dao.UserDAO, logger *zap.Logger) UserService {
	return &UserServiceImpl{
		userDAO: userDAO,
		logger:  logger.With(zap.String("type", "NewUserServiceImpl")),
	}
}

// Login 登录
func (s *UserServiceImpl) Login(ctx context.Context, email, password string) (*responses.UserInfo, error) {
	user, err := s.userDAO.SelectByEmail(email)
	if err != nil {
		s.logger.Error("邮箱不存在", zap.Error(err))
		return nil, err
	}
	if user.Password == password {
		return &responses.UserInfo{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		}, nil
	}
	return nil, ErrPassword
}

// Register 注册
func (s UserServiceImpl) Register(ctx context.Context, vo *responses.RegisterUser) (*responses.UserInfo, error) {
	existUser, err := s.userDAO.SelectByEmail(vo.Email)
	if existUser != nil {
		s.logger.Error("查询邮箱失败", zap.Error(ErrUserExisted))
		return &responses.UserInfo{}, ErrUserExisted
	}
	if err == gorm.ErrRecordNotFound || err == nil {
		newUser := &models.User{
			Username: vo.Username,
			Password: vo.Password,
			Email:    vo.Email,
		}
		err = s.userDAO.Save(newUser)
		if err != nil {
			s.logger.Error("用户创建失败", zap.Error(ErrRegistering))
			return nil, err
		}
		return &responses.UserInfo{
			ID:       newUser.ID,
			Username: newUser.Username,
			Email:    newUser.Email,
		}, nil
	}

	return nil, err
}
