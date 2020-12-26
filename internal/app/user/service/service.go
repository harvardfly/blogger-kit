package service

import (
	"context"
	"errors"
	"time"

	"pkg.zpf.com/golang/blogger-kit/internal/app/user/dao"
	"pkg.zpf.com/golang/blogger-kit/internal/pkg/models"
	"pkg.zpf.com/golang/blogger-kit/internal/pkg/requests"
	"pkg.zpf.com/golang/blogger-kit/internal/pkg/responses"
	"pkg.zpf.com/golang/blogger-kit/internal/pkg/utils/middlewareutil"
	pb "pkg.zpf.com/golang/blogger-kit/protos/user"

	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

var (
	// ErrUserExisted 用户已存在
	ErrUserExisted = errors.New("user is existed")
	// ErrPassword 邮箱与密码不匹配
	ErrPassword = errors.New("email and password are not match")
	// ErrRegistering 邮箱已存在
	ErrRegistering = errors.New("email is registering")
	// ErrAccessToken token验证失败
	ErrAccessToken = errors.New("生成签名错误")
)

// UserService 用户服务
type UserService interface {
	// Login 登录接口
	Login(ctx context.Context, email, password string) (*responses.LoginResponse, error)
	// Register 注册接口
	Register(ctx context.Context, vo *responses.RegisterUser) (*responses.UserInfo, error)
	// FindByID 调用rpc通过ID查询用户信息
	FindByID(ctx context.Context, req *requests.FindByIDRequest) (*responses.UserInfo, error)
	// FindByEmail 调用rpc通过Email查询用户信息
	FindByEmail(ctx context.Context, req *requests.FindByEmailRequest) (*responses.UserInfo, error)
}

// UserServiceImpl 初始默认的
type UserServiceImpl struct {
	userDao    dao.UserDao
	logger     *zap.Logger
	userClient pb.UserClient
}

// NewUserServiceImpl 初始化
func NewUserServiceImpl(userDao dao.UserDao, logger *zap.Logger, userClient pb.UserClient) UserService {
	return &UserServiceImpl{
		userDao:    userDao,
		logger:     logger.With(zap.String("type", "NewUserServiceImpl")),
		userClient: userClient,
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
			return nil, ErrAccessToken
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

// FindByID 调用rpc服务通过ID获取用户信息
func (s *UserServiceImpl) FindByID(ctx context.Context, req *requests.FindByIDRequest) (*responses.UserInfo, error) {
	rpcReq := pb.FindByIDRequest{
		Id: int32(req.ID),
	}

	data, err := s.userClient.FindById(ctx, &rpcReq)
	if err != nil {
		s.logger.Error("通过用户ID获取用户调用失败", zap.Error(err))
		return nil, err
	}
	res := &responses.UserInfo{
		ID:       int(data.Id),
		Username: data.Username,
		Email:    data.Email,
	}

	return res, nil
}

// FindByEmail 调用rpc服务通过Email获取用户信息
func (s *UserServiceImpl) FindByEmail(ctx context.Context, req *requests.FindByEmailRequest) (*responses.UserInfo, error) {
	var res *responses.UserInfo
	rpcReq := pb.FindByEmailRequest{
		Email: req.Email,
	}
	data, err := s.userClient.FindByEmail(ctx, &rpcReq)
	if err != nil {
		s.logger.Error("通过Email获取用户调用失败", zap.Error(err))

	}
	res = &responses.UserInfo{
		ID:       int(data.Id),
		Username: data.Username,
		Email:    data.Email,
	}
	return res, nil
}
