package database

import (
	"go-training/common"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func New(config *common.Config) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(config.Dsn), &gorm.Config{})
}
