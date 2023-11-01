package subscriber

import (
	"context"
	"fmt"
	"go-training/internal/component/pubsub"
	restaurantrepo "go-training/internal/modules/restaurant/repository"
)

func DecreaseLikeCountAfterUserUnLikeRestaurant(restaurantRepo restaurantrepo.RestaurantRepository) consumerJob {
	return consumerJob{
		Title: "Decrease like count after user un-like restaurant",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			likeData := message.Data().(HasRestaurantId)
			fmt.Println("Decrease like count after user un-like restaurant", likeData.GetOwnerId())
			return restaurantRepo.DecreaseLikeCount(ctx, likeData.GetRestaurantId())
		},
	}
}
