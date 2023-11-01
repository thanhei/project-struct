package database

import (
	"fmt"
	"go-training/internal/common"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func New(config *common.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Mysql.User, config.Mysql.Password, config.Mysql.Host, config.Mysql.Port, config.Mysql.Database)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
