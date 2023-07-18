package business

import (
	"context"
	"go-training/modules/restaurant/entity"
)

func (biz *business) CreateRestaurant(ctx context.Context, data *entity.RestaurantCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	if err := biz.restaurantRepo.Create(ctx, data); err != nil {
		return err
	}

	return nil
}
