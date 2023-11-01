package business

import (
	"context"
	"go-training/internal/common"
	"go-training/internal/component/pubsub"
	"go-training/internal/modules/restaurantlike/entity"
)

func (biz *business) UnLikeRestaurant(ctx context.Context, userId, restaurantId int) error {
	err := biz.restaurantLikeRepo.Delete(ctx, userId, restaurantId)

	if err != nil {
		return entity.ErrUserCannotUnLikeRestaurant(err)
	}

	_ = biz.pubsub.Publish(ctx, common.TopicUserDislikeRestaurant, pubsub.NewMessage(&entity.Like{
		restaurantId,
		userId,
		nil,
		nil,
	}))
	//go func() {
	//	defer common.AppRecovery()
	//
	//	job := asyncjob.NewJob(func(ctx context.Context) error {
	//		return biz.decStore.DecreaseLikeCount(ctx, restaurantId)
	//	})
	//
	//	job.SetRetryDurations([]time.Duration{time.Second})
	//
	//	_ = asyncjob.NewGroup(true, job).Run(ctx)
	//}()

	return nil
}
