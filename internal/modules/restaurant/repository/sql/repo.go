package sql

import (
	restaurantrepo "go-training/internal/modules/restaurant/repository"

	"gorm.io/gorm"
)

type sqlRepo struct {
	db *gorm.DB
}

var _ restaurantrepo.RestaurantRepository = (*sqlRepo)(nil)

func NewSQLRepo(db *gorm.DB) *sqlRepo {
	return &sqlRepo{db: db}
}
