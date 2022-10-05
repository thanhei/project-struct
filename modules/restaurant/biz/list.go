package restaurantbiz

import (
	"context"
	"go-training/common"
	restaurantmodel "go-training/modules/restaurant/model"
)

type ListRestaurantStore interface {
	ListDataWithCondition(context context.Context, condition map[string]interface{}, filter *restaurantmodel.Filter, paging *common.Paging, moreKeys ...string) ([]restaurantmodel.Restaurant, error)
}

type listRestaurantBiz struct {
	store ListRestaurantStore
}

func NewListRestaurantBiz(store ListRestaurantStore) *listRestaurantBiz {
	return &listRestaurantBiz{store: store}
}

func (biz *listRestaurantBiz) ListRestaurant(ctx context.Context, filter *restaurantmodel.Filter, paging *common.Paging) ([]restaurantmodel.Restaurant, error) {
	result, err := biz.store.ListDataWithCondition(ctx, nil, filter, paging)

	if err != nil {
		return nil, err
	}

	return result, nil
}
