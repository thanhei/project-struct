package sql

import (
	"context"
	"go-training/internal/common"
	"go-training/internal/modules/restaurant/entity"

	"gorm.io/gorm"
)

func (s *sqlRepo) UpdateData(ctx context.Context, id int, data *entity.RestaurantUpdate) error {
	db := s.db
	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func (s *sqlRepo) IncreaseLikeCount(ctx context.Context, id int) error {
	db := s.db
	if err := db.Table(entity.Restaurant{}.TableName()).Where("id = ?", id).Update("liked_count", gorm.Expr("liked_count + ?", 1)).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

func (s *sqlRepo) DecreaseLikeCount(ctx context.Context, id int) error {
	db := s.db
	if err := db.Table(entity.Restaurant{}.TableName()).Where("id = ?", id).Update("liked_count", gorm.Expr("liked_count - ?", 1)).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
