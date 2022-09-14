package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"go-training/common"
	"go-training/component/app_context"
	restaurantbiz "go-training/modules/restaurant/biz"
	restaurantmodel "go-training/modules/restaurant/model"
	restaurantstorage "go-training/modules/restaurant/storage"
	"net/http"
)

func ListRestaurant(appCtx app_context.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		var pagingData common.Paging
		if err := c.ShouldBind(&pagingData); err != nil {
			c.JSON(http.StatusBadRequest, err)
		}
		pagingData.Fulfill()

		var filter restaurantmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusBadRequest, err)
		}

		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewListRestaurantBiz(store)

		result, err := biz.ListRestaurant(c.Request.Context(), &filter, &pagingData)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"result":     result,
			"pagingData": pagingData,
			"filter":     filter,
		})
	}
}

