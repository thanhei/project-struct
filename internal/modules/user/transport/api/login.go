package api

import (
	"go-training/internal/common"
	"go-training/internal/modules/user/entity"

	"github.com/gin-gonic/gin"
)

func (a *Api) Login() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var data entity.UserLogin

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		account, err := a.biz.Login(c.Request.Context(), &data)
		if err != nil {
			panic(err)
		}

		c.JSON(200, common.SimpleSuccessResponse(account))
	}
}
