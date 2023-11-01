package api

import (
	"go-training/internal/common"
	"go-training/internal/modules/restaurant/entity"

	"github.com/gin-gonic/gin"
)

func (a *Api) UpdateRestaurant() gin.HandlerFunc {
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

		if err := a.biz.UpdateRestaurantBiz(c.Request.Context(), int(uid.GetLocalID()), &data); err != nil {
			c.JSON(401, gin.H{
				"error": err,
			})
			return
		}

		c.JSON(200, common.SimpleSuccessResponse(true))
	}
}
