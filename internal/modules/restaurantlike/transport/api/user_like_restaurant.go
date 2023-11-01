package api

import (
	"go-training/internal/common"
	"go-training/internal/modules/restaurantlike/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

//POST /v1/restaurant/:id/like

func (a *api) UserLikeRestaurant() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(err)
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		data := entity.Like{
			RestaurantId: int(uid.GetLocalID()),
			UserId:       requester.GetUserId(),
		}

		err = a.biz.LikeRestaurant(c.Request.Context(), &data)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
