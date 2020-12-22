package dao

import (
	"pkg.zpf.com/golang/kit-scaffold/internal/pkg/databases"
	"pkg.zpf.com/golang/kit-scaffold/internal/pkg/models"

	"go.uber.org/zap"
)

// UserDao dao数据层方法
type UserDao interface {
	SelectByEmail(email string) (*models.User, error)
	Save(user *models.User) error
}

// UserDaoImpl 默认实现
type UserDaoImpl struct {
	logger *zap.Logger
}

// NewUserDaoImpl 初始化
func NewUserDaoImpl(logger *zap.Logger) UserDao {
	return &UserDaoImpl{
		logger: logger.With(zap.String("type", "NewUserDAOImpl")),
	}
}

// SelectByEmail 查询邮箱
func (d *UserDaoImpl) SelectByEmail(email string) (*models.User, error) {
	user := &models.User{}
	err := databases.DB.Where("email = ?", email).First(user).Error
	d.logger.Info(user.Username, zap.String("username", user.Username))
	if err != nil {
		return nil, err
	}
	return user, err
}

// Save 创建用户
func (d *UserDaoImpl) Save(user *models.User) error {
	return databases.DB.Create(user).Error
}
