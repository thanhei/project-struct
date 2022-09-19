package ginrestaurant

import (
	"go-training/common"
	"go-training/component/app_context"
	restaurantbiz "go-training/modules/restaurant/biz"
	restaurantmodel "go-training/modules/restaurant/model"
	restaurantstorage "go-training/modules/restaurant/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateRestaurant(appCtx app_context.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var data restaurantmodel.RestaurantCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrCannotCreateEntity(restaurantmodel.EntityName, err))
			return
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		data.UserId = requester.GetUserId()

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewCreateRestaurantBiz(store)

		if err := biz.CreateRestaurant(c.Request.Context(), &data); err != nil {
			panic(common.ErrCannotCreateEntity(restaurantmodel.EntityName, err))
			return
		}

		data.GenUID(common.DbTypeRestaurant)

		c.JSON(http.StatusOK, common.NewSuccessResponse(data.FakeId.String(), nil, nil))
	}
}
