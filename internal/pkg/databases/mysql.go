package databases

import (
	"blogger-kit/internal/pkg/config"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitMysql(cfg *config.MySQLConfig) (err error) {
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
	DB.SingularTable(true)
	return
}
