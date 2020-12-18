package service

import (
	"blogger-kit/internal/app/user/dao"
	"blogger-kit/internal/pkg/models"
	"blogger-kit/internal/pkg/responses"
	"blogger-kit/internal/pkg/utils/middlewareutil"
	"context"
	"errors"
	"time"

	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

var (
	ErrUserExisted = errors.New("user is existed")
	ErrPassword    = errors.New("email and password are not match")
	ErrRegistering = errors.New("email is registering")
	AccessTokenErr = errors.New("生成签名错误")
)

// UserService 用户服务
type UserService interface {
	// Login 登录接口
	Login(ctx context.Context, email, password string) (*responses.LoginResponse, error)
	// Register 注册接口
	Register(ctx context.Context, vo *responses.RegisterUser) (*responses.UserInfo, error)
}

// UserServiceImpl 初始默认的
type UserServiceImpl struct {
	userDao dao.UserDao
	logger  *zap.Logger
}

// NewUserServiceImpl 初始化
func NewUserServiceImpl(userDao dao.UserDao, logger *zap.Logger) UserService {
	return &UserServiceImpl{
		userDao: userDao,
		logger:  logger.With(zap.String("type", "NewUserServiceImpl")),
	}
}

// Login 登录
func (s *UserServiceImpl) Login(ctx context.Context, email, password string) (*responses.LoginResponse, error) {
	user, err := s.userDao.SelectByEmail(email)
	if err != nil {
		s.logger.Error("邮箱不存在", zap.Error(err))
		return nil, err
	}
	expired := time.Now().Add(7 * 24 * time.Hour).Unix()
	if user.Password == password {
		accessToken, err := middlewareutil.CreateAccessToken(user.Username, expired)
		if err != nil {
			return nil, AccessTokenErr
		}
		return &responses.LoginResponse{
			AccessToken: accessToken,
			ExpireAt:    expired,
			TimeStamp:   time.Now().Unix(),
		}, nil
	}

	return nil, ErrPassword
}

// Register 注册
func (s UserServiceImpl) Register(ctx context.Context, vo *responses.RegisterUser) (*responses.UserInfo, error) {
	existUser, err := s.userDao.SelectByEmail(vo.Email)
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
		err = s.userDao.Save(newUser)
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
