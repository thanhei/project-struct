package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (a *Api) FindRestaurant() gin.HandlerFunc {
	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		result, err := a.biz.FindRestaurant(c.Request.Context(), id)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, result)
	}
}
