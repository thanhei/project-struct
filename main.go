package main

import (
	"github.com/gin-gonic/gin"
	"go-training/component/app_context"
	"go-training/modules/restaurant/transport/ginrestaurant"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func main() {
	dsn := os.Getenv("MYSQL_CONN_STRING")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	mode := os.Getenv("GIN_MODE")
	if mode != "release" {

	}
	db = db.Debug()
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	appCtx := app_context.NewAppContext(db)
	r.GET("/restaurants", ginrestaurant.ListRestaurant(appCtx))
	r.GET("/restaurants/:id", ginrestaurant.FindRestaurant(appCtx))
	r.POST("/restaurants", ginrestaurant.CreateRestaurant(appCtx))
	r.DELETE("/restaurants/:id", ginrestaurant.DeleteRestaurant(appCtx))
	r.Run()
}
