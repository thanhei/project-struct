package repository

import (
	"context"
	"go-training/internal/common"
	"go-training/internal/modules/restaurantlike/entity"
)

type RestaurantLikeRepository interface {
	Create(ctx context.Context, data *entity.Like) error
	GetRestaurantLike(ctx context.Context, ids []int) (map[int]int, error)
	GetUsersLikeRestaurant(ctx context.Context, conditions map[string]interface{}, filter *entity.Filter, paging *common.Paging, moreKeys ...string) ([]common.SimpleUser, error)
	Delete(ctx context.Context, userId, restaurantId int) error
}
