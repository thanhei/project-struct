package ginuser

import (
	"go-training/common"
	"go-training/component/app_context"
	"go-training/component/hasher/md5"
	"go-training/component/tokenprovider/jwt"
	userbiz "go-training/modules/user/biz"
	usermodel "go-training/modules/user/model"
	userstorage "go-training/modules/user/storage"

	"github.com/gin-gonic/gin"
)

func Login(appCtx app_context.AppContext) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var data usermodel.UserLogin

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())
		store := userstorage.NewSQLStore(db)
		md5 := md5.NewMd5Hash()

		biz := userbiz.NewLoginBusiness(store, tokenProvider, md5, 60*60*24*30)
		account, err := biz.Login(c.Request.Context(), &data)
		if err != nil {
			panic(err)
		}

		c.JSON(200, common.SimpleSuccessResponse(account))
	}
}
