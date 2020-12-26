package dao

import (
	"pkg.zpf.com/golang/blogger-kit/internal/pkg/databases"
	"pkg.zpf.com/golang/blogger-kit/internal/pkg/models"

	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

// UserDao dao数据层方法
type UserDao interface {
	FindByEmail(email string) (models.User, error)
	FindByID(id int32) (models.User, error)
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

// withID 通过单个ID查询对应记录
func withID(typeid int32) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if typeid > 0 {
			db = db.Where("id = ?", typeid)
		}
		return db
	}
}

// FindByEmail 数据库实现根据email获取用户信息
func (r *UserDaoImpl) FindByEmail(email string) (models.User, error) {
	var res models.User
	if err := databases.DB.Model(&models.User{}).Where("email=?", email).First(&res).Error; err != nil {
		r.logger.Error("查询email失败", zap.Error(err))
		return res, err
	}
	return res, nil
}

// FindByID 数据库实现根据ID获取用户信息
func (r *UserDaoImpl) FindByID(id int32) (models.User, error) {
	var res models.User
	if err := databases.DB.Model(&models.User{}).Scopes(withID(id)).First(&res).Error; err != nil {
		r.logger.Error("查询用户ID失败", zap.Error(err))
		return res, err
	}
	return res, nil
}
