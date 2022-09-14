package ginrestaurant

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-training/component/app_context"
	restaurantbiz "go-training/modules/restaurant/biz"
	restaurantmodel "go-training/modules/restaurant/model"
	restaurantstorage "go-training/modules/restaurant/storage"
	"net/http"
	"strconv"
)

func DeleteRestaurant(appCtx app_context.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		db := appCtx.GetMainDBConnection()
		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewDeleteRestaurantBiz(store)

		if err := biz.DeleteRestaurant(c.Request.Context(), id); err != nil {
			if errors.Is(err, restaurantmodel.ErrNotFound) {
				c.JSON(http.StatusNotFound, gin.H{
					"error": err.Error(),
				})
				return
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": true,
		})
	}
}
