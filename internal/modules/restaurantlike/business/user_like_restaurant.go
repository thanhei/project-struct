package business

import (
	"context"
	"go-training/internal/common"
	"go-training/internal/component/pubsub"
	"go-training/internal/modules/restaurantlike/entity"
)

func (biz *business) LikeRestaurant(ctx context.Context, data *entity.Like) error {
	err := biz.restaurantLikeRepo.Create(ctx, data)

	if err != nil {
		return entity.ErrUserCannotLikeRestaurant(err)
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
