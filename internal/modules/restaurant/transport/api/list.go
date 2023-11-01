package api

import (
	"go-training/internal/common"
	"go-training/internal/modules/restaurant/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *Api) ListRestaurant() gin.HandlerFunc {
	return func(c *gin.Context) {
		var pagingData common.Paging
		if err := c.ShouldBind(&pagingData); err != nil {
			c.JSON(http.StatusBadRequest, err)
		}
		pagingData.Fulfill()

		var filter entity.Filter
		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusBadRequest, err)
		}

		result, err := a.biz.ListRestaurant(c.Request.Context(), &filter, &pagingData)

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
