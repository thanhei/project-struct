package api

import (
	"go-training/internal/common"

	restanrantlikebiz "go-training/internal/modules/restaurantlike/business"

	"github.com/gin-gonic/gin"
)

type api struct {
	biz restanrantlikebiz.RestaurantLikeBusiness
}

func NewApi(biz restanrantlikebiz.RestaurantLikeBusiness) common.AppModule {
	return &api{
		biz: biz,
	}
}

func (a *api) SetupRoutes(r *gin.RouterGroup) {
	r.GET("/:id/liked-users", a.ListUserLikeRestaurant())
	r.POST("/:id/like", a.UserLikeRestaurant())
	r.DELETE("/:id/unlike", a.UserUnLikeRestaurant())
}
