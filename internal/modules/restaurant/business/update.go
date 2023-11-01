package business

import (
	"context"
	"errors"
	"go-training/internal/common"
	"go-training/internal/modules/restaurant/entity"
)

func (biz *business) UpdateRestaurantBiz(ctx context.Context, id int, data *entity.RestaurantUpdate) error {
	oldData, err := biz.restaurantRepo.FindDataWithCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		if err != common.RecordNotFound {
			return common.RecordNotFound
		}
		return err
	}
	if oldData.Status == 0 {
		return errors.New("Data deleted")
	}

	if err := biz.restaurantRepo.UpdateData(ctx, id, data); err != nil {
		return err
	}
	return nil
}
