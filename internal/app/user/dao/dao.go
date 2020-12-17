package dao

import (
	"blogger-kit/internal/pkg/databases"
	"blogger-kit/internal/pkg/models"

	"go.uber.org/zap"
)

type UserDAO interface {
	SelectByEmail(email string) (*models.User, error)
	Save(user *models.User) error
}

type UserDAOImpl struct {
	logger *zap.Logger
}

// NewUserDAOImpl 初始化
func NewUserDAOImpl(logger *zap.Logger) UserDAO {
	return &UserDAOImpl{
		logger: logger.With(zap.String("type", "NewUserDAOImpl")),
	}
}

// SelectByEmail 查询邮箱
func (d *UserDAOImpl) SelectByEmail(email string) (*models.User, error) {
	user := &models.User{}
	err := databases.DB.Where("email = ?", email).First(user).Error
	d.logger.Info(user.Username, zap.String("username", user.Username))
	return user, err
}

// Save 创建用户
func (d *UserDAOImpl) Save(user *models.User) error {
	return databases.DB.Create(user).Error
}
