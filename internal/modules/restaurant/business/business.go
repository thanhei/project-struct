package business

import (
	"context"
	"go-training/internal/common"
	"go-training/internal/modules/restaurant/entity"
	restaurantrepo "go-training/internal/modules/restaurant/repository"
)

type RestaurantBusiness interface {
	CreateRestaurant(ctx context.Context, data *entity.RestaurantCreate) error
	DeleteRestaurant(ctx context.Context, id int) error
	FindRestaurant(ctx context.Context, id int) (*entity.Restaurant, error)
	ListRestaurant(ctx context.Context, filter *entity.Filter, paging *common.Paging) ([]entity.Restaurant, error)
	UpdateRestaurantBiz(ctx context.Context, id int, data *entity.RestaurantUpdate) error
}

type business struct {
	restaurantRepo restaurantrepo.RestaurantRepository
}

var _ RestaurantBusiness = (*business)(nil)

func NewBusiness(restaurantRepo restaurantrepo.RestaurantRepository) *business {
	return &business{restaurantRepo: restaurantRepo}
}
