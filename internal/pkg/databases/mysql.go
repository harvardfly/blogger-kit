package databases

import (
	"blogger-kit/internal/pkg/config"
	"fmt"
	"log"
	"time"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

// DB 定义gorm初始变量
var DB *gorm.DB

// InitMysql 初始化mysql连接
func InitMysql(cfg *config.MySQLConfig) (err error) {
	// 连接MySQL驱动
	DB, err = gorm.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DB,
		),
	)
	if err != nil {
		log.Println(err)
		return
	}
	// Ping MySQL
	err = DB.DB().Ping()
	if err != nil {
		return errors.Wrap(err, "mysql ping fail")
	}
	// Debug模式下输出sql信息
	if cfg.Debug {
		DB = DB.Debug()
	}
	DB.DB().SetConnMaxLifetime(time.Minute * 10)
	DB.SingularTable(true)
	return
}
