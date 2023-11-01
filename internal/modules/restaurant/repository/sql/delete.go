package sql

import (
	"context"
	"go-training/internal/modules/restaurant/entity"
)

func (s *sqlRepo) Delete(context context.Context, id int) error {
	if err := s.db.Table(entity.Restaurant{}.TableName()).Where("id = ?", id).Updates(map[string]interface{}{
		"status": 0,
	}).Error; err != nil {
		return err
	}

	return nil
}
