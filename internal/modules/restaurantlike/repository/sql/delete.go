package restaurantlikereposql

import (
	"context"
	"go-training/internal/common"
	"go-training/internal/modules/restaurantlike/entity"
)

func (s *sqlRepo) Delete(ctx context.Context, userId, restaurantId int) error {
	db := s.db

	if err := db.Table(entity.Like{}.TableName()).Where("user_id = ? and restaurant_id = ?", userId, restaurantId).Delete(nil).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
