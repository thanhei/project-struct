package ginrestaurantlike

import (
	"go-training/common"
	"go-training/component/app_context"
	rstlikebiz "go-training/modules/restaurantlike/biz"
	restaurantlikemodel "go-training/modules/restaurantlike/model"
	restaurantlikestorage "go-training/modules/restaurantlike/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

//GET /v1/restaurant/:id/liked-user

func ListUserLikeRestaurant(appCtx app_context.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		//var filter restaurantlikemodel.Filter
		//

		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}

		filter := restaurantlikemodel.Filter{
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

		store := restaurantlikestorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := rstlikebiz.NewListUserLikeRestaurant(store)

		result, err := biz.ListUser(c.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
