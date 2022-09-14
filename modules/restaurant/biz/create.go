package restaurantbiz

import (
	"context"
	restaurantmodel "go-training/modules/restaurant/model"
)

type CreateRestaurantStore interface {
	Create(context context.Context, data *restaurantmodel.RestaurantCreate) error
}

type createRestaurantBiz struct {
	store CreateRestaurantStore
}

func NewCreateRestaurantBiz(store CreateRestaurantStore) *createRestaurantBiz {
	return &createRestaurantBiz{store: store}
}

func (biz *createRestaurantBiz) CreateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {

	// Validate data theo Biz
	// Ten khong duoc trung, Name unique
	// Validate dia chi co ton tai hay khong

	if err := data.Validate(); err != nil {
		return err
	}

	if err := biz.store.Create(ctx, data); err != nil {
		return err
	}

	return nil
}
