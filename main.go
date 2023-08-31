package main

import (
	"context"
	"fmt"
	"go-training/common"
	"go-training/component/app_context"
	"go-training/component/database"
	"go-training/memcache"
	"go-training/middleware"
	"go-training/modules/restaurant/transport/ginrestaurant"
	"go-training/modules/restaurantlike/transport/ginrestaurantlike"
	"go-training/modules/user/transport/ginuser"
	"go-training/pubsub"
	"go-training/pubsub/pblocal"
	"go-training/subscriber"
	"log"

	userstorage "go-training/modules/user/storage"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			common.NewConfig,
			database.New,
			fx.Annotate(pblocal.NewPubSub, fx.As(new(pubsub.Pubsub))),
			fx.Annotate(app_context.NewAppContext, fx.As(new(app_context.AppContext))),
			subscriber.NewEngine,
		),
		fx.Invoke(registerHooks),
	).Run()
}

func registerHooks(lifecycle fx.Lifecycle, appCtx app_context.AppContext, subscriber *subscriber.ConsumerEngine) {
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			r := gin.Default()

			if err := subscriber.Start(); err != nil {
				log.Fatalln(err)
			}

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
			return r.Run()
		},
		OnStop: func(context.Context) error {
			fmt.Println("stop")
			db, err := appCtx.GetMainDBConnection().DB()
			if err != nil {
				return err
			}
			return db.Close()
		},
	})
}
