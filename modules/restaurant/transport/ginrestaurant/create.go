package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"go-training/component/app_context"
	restaurantbiz "go-training/modules/restaurant/biz"
	restaurantmodel "go-training/modules/restaurant/model"
	restaurantstorage "go-training/modules/restaurant/storage"
	"net/http"
)

func CreateRestaurant(appCtx app_context.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var data restaurantmodel.RestaurantCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		db := appCtx.GetMainDBConnection()
		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewCreateRestaurantBiz(store)

		if err := biz.CreateRestaurant(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
}
