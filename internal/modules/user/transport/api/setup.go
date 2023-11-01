package api

import (
	"go-training/internal/common"
	userbiz "go-training/internal/modules/user/business"

	"github.com/gin-gonic/gin"
)

type Api struct {
	biz userbiz.UserBusiness
}

func NewApi(biz userbiz.UserBusiness) common.AppModule {
	return &Api{biz}
}

func (a *Api) SetupRoutes(r *gin.RouterGroup) {
	r.POST("/register", a.Register())
	r.POST("/login", a.Login())
}
