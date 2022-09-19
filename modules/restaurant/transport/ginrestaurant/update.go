package ginrestaurant

import (
	"go-training/common"
	"go-training/component/app_context"
	restaurantbiz "go-training/modules/restaurant/biz"
	restaurantmodel "go-training/modules/restaurant/model"
	restaurantstorage "go-training/modules/restaurant/storage"

	"github.com/gin-gonic/gin"
)

func UpdateRestaurant(appCtx app_context.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		//id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(401, gin.H{
				"error": err,
			})
			return
		}

		var data restaurantmodel.RestaurantUpdate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(401, gin.H{
				"error": err,
			})
			return
		}

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewUpdateRestaurantBiz(store)

		if err := biz.UpdateRestaurantBiz(c.Request.Context(), int(uid.GetLocalID()), &data); err != nil {
			c.JSON(401, gin.H{
				"error": err,
			})
			return
		}

		c.JSON(200, common.SimpleSuccessResponse(true))
	}
}
