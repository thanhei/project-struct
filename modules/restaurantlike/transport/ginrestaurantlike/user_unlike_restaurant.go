package ginrestaurantlike

import (
	"go-training/common"
	"go-training/component/app_context"
	rstlikebiz "go-training/modules/restaurantlike/biz"
	restaurantlikestorage "go-training/modules/restaurantlike/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

//PUT /v1/restaurant/:id/unlike

func UserUnLikeRestaurant(appCtx app_context.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(err)
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		store := restaurantlikestorage.NewSQLStore(appCtx.GetMainDBConnection())
		//decStore := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		pubsub := appCtx.GetPubsub()
		biz := rstlikebiz.NewUserUnLikeRestaurantBiz(store, pubsub)

		err = biz.UnLikeRestaurant(c.Request.Context(), requester.GetUserId(), int(uid.GetLocalID()))
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
