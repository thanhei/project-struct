package business

import (
	"context"
	"go-training/internal/common"
	"go-training/internal/modules/restaurant/entity"
)

func (biz *business) DeleteRestaurant(ctx context.Context, id int) error {
	oldData, err := biz.restaurantRepo.FindDataWithCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		if err != common.RecordNotFound {
			return common.RecordNotFound
		}
		return err
	}
	if oldData.Status == 0 {
		return common.ErrEntityDeleted(entity.EntityName, err)
	}

	if err := biz.restaurantRepo.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}
