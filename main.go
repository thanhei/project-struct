package main

import (
	"go-training/component/app_context"
	"go-training/memcache"
	"go-training/middleware"
	"go-training/modules/restaurant/transport/ginrestaurant"
	"go-training/modules/restaurantlike/transport/ginrestaurantlike"
	userstorage "go-training/modules/user/storage"
	"go-training/modules/user/transport/ginuser"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := os.Getenv("MYSQL_CONN_STRING")
	secretKey := os.Getenv("SYSTEM_SECRET")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	mode := os.Getenv("GIN_MODE")
	if mode != "release" {

	}
	db = db.Debug()
	if err != nil {
		panic(err)
	}
	appCtx := app_context.NewAppContext(db, secretKey)

	r := gin.Default()

	r.Use(middleware.Recover(appCtx))

	v1 := r.Group("/v1")
	v1.POST("/register", ginuser.Register(appCtx))
	v1.POST("/login", ginuser.Login(appCtx))

	userStore := userstorage.NewSQLStore(appCtx.GetMainDBConnection())
	userCaching := memcache.NewUserCaching(memcache.NewCaching(), userStore)

	restaurants := v1.Group("/restaurants", middleware.RequireAuth(appCtx, userCaching))
	{
		restaurants.POST("", ginrestaurant.CreateRestaurant(appCtx))
		restaurants.GET("", ginrestaurant.ListRestaurant(appCtx))
		restaurants.PATCH("/:id", ginrestaurant.UpdateRestaurant(appCtx))
		restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))

		restaurants.GET("/:id/liked-users", ginrestaurantlike.ListUserLikeRestaurant(appCtx))

		restaurants.POST("/:id/like", ginrestaurantlike.UserLikeRestaurant(appCtx))
		restaurants.DELETE("/:id/unlike", ginrestaurantlike.UserUnLikeRestaurant(appCtx))
	}
	r.Run()
}
