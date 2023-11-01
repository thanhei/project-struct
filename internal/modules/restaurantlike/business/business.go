package business

import (
	"context"
	"go-training/internal/common"
	"go-training/internal/component/pubsub"
	"go-training/internal/modules/restaurantlike/entity"
	restaurantlikerepo "go-training/internal/modules/restaurantlike/repository"
)

type RestaurantLikeBusiness interface {
	LikeRestaurant(ctx context.Context, data *entity.Like) error
	ListUser(ctx context.Context, filter *entity.Filter, paging *common.Paging) ([]common.SimpleUser, error)
	UnLikeRestaurant(ctx context.Context, userId, restaurantId int) error
}

type business struct {
	restaurantLikeRepo restaurantlikerepo.RestaurantLikeRepository
	pubsub             pubsub.Pubsub
}

var _ RestaurantLikeBusiness = (*business)(nil)

func NewBusiness(restaurantLikeRepo restaurantlikerepo.RestaurantLikeRepository, pubsub pubsub.Pubsub) *business {
	return &business{
		restaurantLikeRepo: restaurantLikeRepo,
		pubsub:             pubsub,
	}
}
