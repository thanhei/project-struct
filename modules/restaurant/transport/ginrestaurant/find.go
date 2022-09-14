package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"go-training/component/app_context"
	restaurantbiz "go-training/modules/restaurant/biz"
	restaurantstorage "go-training/modules/restaurant/storage"
	"net/http"
	"strconv"
)

func FindRestaurant(appCtx app_context.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewFindRestaurantBiz(store)

		result, err := biz.FindRestaurant(c.Request.Context(), id)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, result)
	}
}
