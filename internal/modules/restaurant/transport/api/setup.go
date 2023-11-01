package api

import (
	"go-training/internal/common"

	restaurantbiz "go-training/internal/modules/restaurant/business"

	"github.com/gin-gonic/gin"
)

type Api struct {
	biz restaurantbiz.RestaurantBusiness
}

func NewApi(biz restaurantbiz.RestaurantBusiness) common.AppModule {
	return &Api{biz}
}

func (a *Api) SetupRoutes(r *gin.RouterGroup) {
	r.POST("", a.CreateRestaurant())
	r.GET("", a.ListRestaurant())
	r.PATCH("/:id", a.UpdateRestaurant())
	r.DELETE("/:id", a.DeleteRestaurant())
}
