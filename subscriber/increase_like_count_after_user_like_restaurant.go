package subscriber

import (
	"context"
	"go-training/common"
	"go-training/component/app_context"
	restaurantstorage "go-training/modules/restaurant/storage"
	"go-training/pubsub"
	"log"
)

type HasRestaurantId interface {
	GetRestaurantId() int
	GetOwnerId() int
}

// run with setup without lib
func IncreaseLikeCountAfterUserLikeRestaurant(appCtx app_context.AppContext, ctx context.Context) {
	c, _ := appCtx.GetPubsub().Subscribe(ctx, common.TopicUserLikeRestaurant)

	store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())

	go func() {
		defer common.AppRecovery()
		msg := <-c
		likeData := msg.Data().(HasRestaurantId)

		_ = store.IncreaseLikeCount(ctx, likeData.GetRestaurantId())
	}()
}

// have engine
func RunIncreaseLikeCountAfterUserLikeRestaurant(appCtx app_context.AppContext) consumerJob {
	return consumerJob{
		Title: "Increase like count after user liked restaurant",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
			likeData := message.Data().(HasRestaurantId)
			return store.IncreaseLikeCount(ctx, likeData.GetRestaurantId())
		},
	}
}

func NotificationAfterUserLikeRestaurant(appCtx app_context.AppContext) consumerJob {

	return consumerJob{
		Title: "Notification after user liked restaurant",
		Hld: func(ctx context.Context, message *pubsub.Message) error {

			likeData := message.Data().(HasRestaurantId)
			// do something
			log.Println("Notification after user liked restaurant", likeData.GetOwnerId())
			return nil
		},
	}
}
