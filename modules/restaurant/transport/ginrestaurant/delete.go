package ginrestaurant

import (
	"go-training/common"
	"go-training/component/app_context"
	"go-training/modules/restaurant/business"
	"go-training/modules/restaurant/repository/sql"

	"github.com/gin-gonic/gin"
)

func DeleteRestaurant(appCtx app_context.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		//id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(401, gin.H{
				"error": err,
			})
			return
		}

		store := sql.NewSQLRepo(appCtx.GetMainDBConnection())
		biz := business.NewBusiness(store)

		if err := biz.DeleteRestaurant(c.Request.Context(), int(uid.GetLocalID())); err != nil {
			c.JSON(401, gin.H{
				"error": err,
			})
			return
		}

		c.JSON(200, common.SimpleSuccessResponse(true))
	}
}
