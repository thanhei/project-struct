package api

import (
	"go-training/internal/common"
	"go-training/internal/modules/restaurantlike/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

//GET /v1/restaurant/:id/liked-user

func (a *api) ListUserLikeRestaurant() gin.HandlerFunc {
	return func(c *gin.Context) {

		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}

		filter := entity.Filter{
			RestaurantId: int(uid.GetLocalID()),
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}

		paging.Fulfill()

		result, err := a.biz.ListUser(c.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
