package rstlikebiz

import (
	"context"
	"go-training/common"
	restaurantlikemodel "go-training/modules/restaurantlike/model"
	"go-training/pubsub"
)

type UserLikeRestaurantStore interface {
	Create(ctx context.Context, data *restaurantlikemodel.Like) error
}

type IncreaseLikeCountStore interface {
	IncreaseLikeCount(ctx context.Context, id int) error
}

type userLikeRestaurantBiz struct {
	store UserLikeRestaurantStore
	// incStore IncreaseLikeCountStore
	pubsub pubsub.Pubsub
}

func NewUserLikeRestaurantBiz(
	store UserLikeRestaurantStore,
	// incStore IncreaseLikeCountStore,
	pubsub pubsub.Pubsub,
) *userLikeRestaurantBiz {
	return &userLikeRestaurantBiz{
		store:  store,
		pubsub: pubsub,
		// incStore: incStore,
	}
}

func (biz *userLikeRestaurantBiz) LikeRestaurant(ctx context.Context, data *restaurantlikemodel.Like) error {
	err := biz.store.Create(ctx, data)

	if err != nil {
		return restaurantlikemodel.ErrUserCannotLikeRestaurant(err)
	}

	biz.pubsub.Publish(ctx, common.TopicUserLikeRestaurant, pubsub.NewMessage(data))

	// go func() {
	// 	defer common.AppRecovery()
	// 	job := asyncjob.NewJob(func(ctx context.Context) error {
	// 		return biz.incStore.IncreaseLikeCount(ctx, data.RestaurantId)
	// 	})

	// 	_ = asyncjob.NewGroup(true, job).Run(ctx)
	// }()

	return nil
}
