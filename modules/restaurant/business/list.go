package business

import (
	"context"
	"go-training/common"
	"go-training/modules/restaurant/entity"
)

func (biz *business) ListRestaurant(ctx context.Context, filter *entity.Filter, paging *common.Paging) ([]entity.Restaurant, error) {
	result, err := biz.restaurantRepo.ListDataWithCondition(ctx, nil, filter, paging)

	if err != nil {
		return nil, err
	}

	return result, nil
}
