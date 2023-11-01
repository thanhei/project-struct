package restaurantlikereposql

import (
	restaurantlikerepo "go-training/internal/modules/restaurantlike/repository"

	"gorm.io/gorm"
)

type sqlRepo struct {
	db *gorm.DB
}

var _ restaurantlikerepo.RestaurantLikeRepository = (*sqlRepo)(nil)

func NewSQLRepo(db *gorm.DB) *sqlRepo {
	return &sqlRepo{db: db}
}
