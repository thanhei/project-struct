package api

import (
	"go-training/internal/common"
	"go-training/internal/modules/user/entity"

	"github.com/gin-gonic/gin"
)

func (a *Api) Register() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var data entity.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		if err := a.biz.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(200, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
