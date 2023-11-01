package sql

import (
	"context"
	"go-training/internal/modules/restaurant/entity"
)

func (s *sqlRepo) Create(context context.Context, data *entity.RestaurantCreate) error {
	if err := s.db.Create(data).Error; err != nil {
		return err
	}

	return nil
}
