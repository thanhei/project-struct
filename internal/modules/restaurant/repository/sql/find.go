package sql

import (
	"context"
	"errors"
	"go-training/internal/modules/restaurant/entity"

	"gorm.io/gorm"
)

func (s *sqlRepo) FindDataWithCondition(context context.Context, condition map[string]interface{}, moreKeys ...string) (*entity.Restaurant, error) {
	var data entity.Restaurant

	if err := s.db.Where(condition).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("record not found")
		}
		return nil, err
	}

	return &data, nil
}
