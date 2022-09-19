package subscriber

import (
	"context"
	"fmt"
	"go-training/component/app_context"
	restaurantstorage "go-training/modules/restaurant/storage"
	"go-training/pubsub"
)

func DecreaseLikeCountAfterUserUnLikeRestaurant(appCtx app_context.AppContext) consumerJob {
	return consumerJob{
		Title: "Decrease like count after user un-like restaurant",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
			likeData := message.Data().(HasRestaurantId)
			fmt.Println("Decrease like count after user un-like restaurant", likeData.GetOwnerId())
			return store.DecreaseLikeCount(ctx, likeData.GetRestaurantId())
		},
	}
}
