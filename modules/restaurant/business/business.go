package business

import (
	"context"
	"go-training/common"
	"go-training/modules/restaurant/entity"
)

type RestaurantRepository interface {
	Create(context context.Context, data *entity.RestaurantCreate) error
	Delete(context context.Context, id int) error
	FindDataWithCondition(context context.Context, condition map[string]interface{}, moreKeys ...string) (*entity.Restaurant, error)
	ListDataWithCondition(context context.Context, condition map[string]interface{}, filter *entity.Filter, paging *common.Paging, moreKeys ...string) ([]entity.Restaurant, error)
	UpdateData(ctx context.Context, id int, data *entity.RestaurantUpdate) error
}

type business struct {
	restaurantRepo RestaurantRepository
}

func NewBusiness(restaurantRepo RestaurantRepository) *business {
	return &business{restaurantRepo: restaurantRepo}
}
