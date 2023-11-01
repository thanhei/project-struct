package repository

import (
	"context"
	"go-training/internal/common"
	"go-training/internal/modules/restaurant/entity"
)

type RestaurantRepository interface {
	Create(context context.Context, data *entity.RestaurantCreate) error
	Delete(context context.Context, id int) error
	FindDataWithCondition(context context.Context, condition map[string]interface{}, moreKeys ...string) (*entity.Restaurant, error)
	ListDataWithCondition(context context.Context, condition map[string]interface{}, filter *entity.Filter, paging *common.Paging, moreKeys ...string) ([]entity.Restaurant, error)
	UpdateData(ctx context.Context, id int, data *entity.RestaurantUpdate) error
	IncreaseLikeCount(ctx context.Context, id int) error
	DecreaseLikeCount(ctx context.Context, id int) error
}
