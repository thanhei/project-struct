package restaurantstorage

import (
	"context"
	"errors"
	restaurantmodel "go-training/modules/restaurant/model"
	"gorm.io/gorm"
)

func (s *sqlStore) FindDataWithCondition(context context.Context, condition map[string]interface{}, moreKeys ...string) (*restaurantmodel.Restaurant, error) {
	var data restaurantmodel.Restaurant

	if err := s.db.Where(condition).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("record not found")
		}
		return nil, err
	}

	return &data, nil
}
