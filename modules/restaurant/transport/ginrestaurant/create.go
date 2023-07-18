package ginrestaurant

import (
	"go-training/common"
	"go-training/component/app_context"
	"go-training/modules/restaurant/business"
	"go-training/modules/restaurant/entity"
	"go-training/modules/restaurant/repository/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateRestaurant(appCtx app_context.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var data entity.RestaurantCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrCannotCreateEntity(entity.EntityName, err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		data.UserId = requester.GetUserId()

		store := sql.NewSQLRepo(appCtx.GetMainDBConnection())
		biz := business.NewBusiness(store)

		if err := biz.CreateRestaurant(c.Request.Context(), &data); err != nil {
			panic(common.ErrCannotCreateEntity(entity.EntityName, err))
		}

		data.GenUID(common.DbTypeRestaurant)

		c.JSON(http.StatusOK, common.NewSuccessResponse(data.FakeId.String(), nil, nil))
	}
}
