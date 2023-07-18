package ginrestaurant

import (
	"go-training/component/app_context"
	"go-training/modules/restaurant/business"
	"go-training/modules/restaurant/repository/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FindRestaurant(appCtx app_context.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		store := sql.NewSQLRepo(db)
		biz := business.NewBusiness(store)

		result, err := biz.FindRestaurant(c.Request.Context(), id)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, result)
	}
}
