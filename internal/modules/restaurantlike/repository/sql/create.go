package restaurantlikereposql

import (
	"context"
	"go-training/internal/common"
	"go-training/internal/modules/restaurantlike/entity"
)

func (s *sqlRepo) Create(ctx context.Context, data *entity.Like) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
