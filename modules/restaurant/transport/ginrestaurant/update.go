package ginrestaurant

import (
	"go-training/common"
	"go-training/component/app_context"
	"go-training/modules/restaurant/business"
	"go-training/modules/restaurant/entity"
	"go-training/modules/restaurant/repository/sql"

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

		var data entity.RestaurantUpdate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(401, gin.H{
				"error": err,
			})
			return
		}

		store := sql.NewSQLRepo(appCtx.GetMainDBConnection())
		biz := business.NewBusiness(store)

		if err := biz.UpdateRestaurantBiz(c.Request.Context(), int(uid.GetLocalID()), &data); err != nil {
			c.JSON(401, gin.H{
				"error": err,
			})
			return
		}

		c.JSON(200, common.SimpleSuccessResponse(true))
	}
}
