package rstlikebiz

import (
	"context"
	"go-training/common"
	"go-training/component/asyncjob"
	restaurantlikemodel "go-training/modules/restaurantlike/model"
	"time"
)

type UserUnLikeRestaurantStore interface {
	Delete(ctx context.Context, userId, restaurantId int) error
}

//
type DecreaseLikeCountStore interface {
	DecreaseLikeCount(ctx context.Context, id int) error
}

type userUnLikeRestaurantBiz struct {
	store    UserUnLikeRestaurantStore
	decStore DecreaseLikeCountStore
}

func NewUserUnLikeRestaurantBiz(store UserUnLikeRestaurantStore, decStore DecreaseLikeCountStore) *userUnLikeRestaurantBiz {
	return &userUnLikeRestaurantBiz{
		store:    store,
		decStore: decStore,
	}
}

func (biz *userUnLikeRestaurantBiz) UnLikeRestaurant(ctx context.Context, userId, restaurantId int) error {
	err := biz.store.Delete(ctx, userId, restaurantId)

	if err != nil {
		return restaurantlikemodel.ErrUserCannotUnLikeRestaurant(err)
	}

	go func() {
		defer common.AppRecovery()

		job := asyncjob.NewJob(func(ctx context.Context) error {
			return biz.decStore.DecreaseLikeCount(ctx, restaurantId)
		})

		job.SetRetryDurations([]time.Duration{time.Second})

		_ = asyncjob.NewGroup(true, job).Run(ctx)
	}()

	return nil
}
