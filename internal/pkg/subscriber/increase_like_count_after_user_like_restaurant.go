package subscriber

import (
	"context"
	"go-training/internal/component/pubsub"
	restaurantrepo "go-training/internal/modules/restaurant/repository"
	"log"
)

type HasRestaurantId interface {
	GetRestaurantId() int
	GetOwnerId() int
}

func RunIncreaseLikeCountAfterUserLikeRestaurant(restaurantRepo restaurantrepo.RestaurantRepository) consumerJob {
	return consumerJob{
		Title: "Increase like count after user liked restaurant",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			likeData := message.Data().(HasRestaurantId)
			return restaurantRepo.IncreaseLikeCount(ctx, likeData.GetRestaurantId())
		},
	}
}

func NotificationAfterUserLikeRestaurant() consumerJob {
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
