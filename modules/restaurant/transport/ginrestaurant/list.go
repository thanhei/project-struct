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

func ListRestaurant(appCtx app_context.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		var pagingData common.Paging
		if err := c.ShouldBind(&pagingData); err != nil {
			c.JSON(http.StatusBadRequest, err)
		}
		pagingData.Fulfill()

		var filter entity.Filter
		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusBadRequest, err)
		}

		store := sql.NewSQLRepo(db)
		biz := business.NewBusiness(store)

		result, err := biz.ListRestaurant(c.Request.Context(), &filter, &pagingData)

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)

			if i == len(result)-1 {
				pagingData.NextCursor = result[i].FakeId.String()
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"result":     result,
			"pagingData": pagingData,
			"filter":     filter,
		})
	}
}
