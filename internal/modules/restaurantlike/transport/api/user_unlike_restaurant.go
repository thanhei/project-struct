package api

import (
	"go-training/internal/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

//PUT /v1/restaurant/:id/unlike

func (a *api) UserUnLikeRestaurant() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(err)
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		err = a.biz.UnLikeRestaurant(c.Request.Context(), requester.GetUserId(), int(uid.GetLocalID()))
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
