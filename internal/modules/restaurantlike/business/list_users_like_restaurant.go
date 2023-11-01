package business

import (
	"context"
	"go-training/internal/common"
	"go-training/internal/modules/restaurantlike/entity"
)

func (biz *business) ListUser(ctx context.Context, filter *entity.Filter, paging *common.Paging) ([]common.SimpleUser, error) {
	users, err := biz.restaurantLikeRepo.GetUsersLikeRestaurant(ctx, nil, filter, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(entity.EntityName, err)
	}

	return users, nil
}
