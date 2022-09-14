package restaurantbiz

import (
	"context"
	"fmt"
	restaurantmodel "go-training/modules/restaurant/model"
)

type DeleteRestaurantStore interface {
	Delete(context context.Context, id int) error
	FindDataWithCondition(context context.Context, condition map[string]interface{}, moreKeys ...string) (*restaurantmodel.Restaurant, error)
}

type deleteRestaurantBiz struct {
	store DeleteRestaurantStore
}

func NewDeleteRestaurantBiz(store DeleteRestaurantStore) *deleteRestaurantBiz {
	return &deleteRestaurantBiz{store: store}
}

func (biz *deleteRestaurantBiz) DeleteRestaurant(ctx context.Context, id int) error {

	// Validate data of Biz layer
	// Find Restaurant by id
	restaurant, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"id": id})
	fmt.Println(restaurant)
	if err != nil {
		return err
	}
	if restaurant.Status == 0 {
		return restaurantmodel.ErrNotFound
	}

	if err := biz.store.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}
