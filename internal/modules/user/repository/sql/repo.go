package userreposql

import (
	userrepo "go-training/internal/modules/user/repository"

	"gorm.io/gorm"
)

type sqlRepo struct {
	db *gorm.DB
}

var _ userrepo.UserRepository = (*sqlRepo)(nil)

func NewSQLRepo(db *gorm.DB) *sqlRepo {
	return &sqlRepo{db: db}
}
