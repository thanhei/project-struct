package api

import (
	"go-training/internal/common"
	"go-training/internal/modules/restaurant/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *Api) CreateRestaurant() gin.HandlerFunc {
	return func(c *gin.Context) {

		var data entity.RestaurantCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrCannotCreateEntity(entity.EntityName, err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		data.UserId = requester.GetUserId()

		if err := a.biz.CreateRestaurant(c.Request.Context(), &data); err != nil {
			panic(common.ErrCannotCreateEntity(entity.EntityName, err))
		}

		data.GenUID(common.DbTypeRestaurant)

		c.JSON(http.StatusOK, common.NewSuccessResponse(data.FakeId.String(), nil, nil))
	}
}
