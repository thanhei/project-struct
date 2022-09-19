package ginuser

import (
	"go-training/common"
	"go-training/component/app_context"
	"go-training/component/hasher/md5"
	userbiz "go-training/modules/user/biz"
	usermodel "go-training/modules/user/model"
	userstorage "go-training/modules/user/storage"

	"github.com/gin-gonic/gin"
)

func Register(appCtx app_context.AppContext) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var data usermodel.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := userstorage.NewSQLStore(db)
		md5 := md5.NewMd5Hash()
		biz := userbiz.NewRegisterStorage(store, md5)

		if err := biz.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(200, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
