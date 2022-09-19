package ginrestaurantlike

import (
	"go-training/common"
	"go-training/component/app_context"
	restaurantstorage "go-training/modules/restaurant/storage"
	rstlikebiz "go-training/modules/restaurantlike/biz"
	restaurantlikemodel "go-training/modules/restaurantlike/model"
	restaurantlikestorage "go-training/modules/restaurantlike/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

//POST /v1/restaurant/:id/like

func UserLikeRestaurant(appCtx app_context.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(err)
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		data := restaurantlikemodel.Like{
			RestaurantId: int(uid.GetLocalID()),
			UserId:       requester.GetUserId(),
		}

		store := restaurantlikestorage.NewSQLStore(appCtx.GetMainDBConnection())
		incStore := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := rstlikebiz.NewUserLikeRestaurantBiz(store, incStore)

		err = biz.LikeRestaurant(c.Request.Context(), &data)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
