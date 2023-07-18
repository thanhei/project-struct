package business

import (
	"context"
	"go-training/modules/restaurant/entity"
)

func (biz *business) FindRestaurant(ctx context.Context, id int) (*entity.Restaurant, error) {
	result, err := biz.restaurantRepo.FindDataWithCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return nil, err
	}

	return result, nil
}
