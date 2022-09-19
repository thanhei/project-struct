package subscriber

import (
	"context"
	"go-training/component/app_context"
)

func Setup(appCtx app_context.AppContext) {
	IncreaseLikeCountAfterUserLikeRestaurant(appCtx, context.Background())
}
