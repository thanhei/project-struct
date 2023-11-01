package api

import (
	"go-training/internal/common"

	"github.com/gin-gonic/gin"
)

func (a *Api) DeleteRestaurant() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		//id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(401, gin.H{
				"error": err,
			})
			return
		}

		if err := a.biz.DeleteRestaurant(c.Request.Context(), int(uid.GetLocalID())); err != nil {
			c.JSON(401, gin.H{
				"error": err,
			})
			return
		}

		c.JSON(200, common.SimpleSuccessResponse(true))
	}
}
