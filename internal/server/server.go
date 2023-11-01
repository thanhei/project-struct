package server

import (
	"go-training/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewServer() *gin.Engine {
	r := gin.Default()
	// r.Use(middleware.Cors("*"))
	r.Use(middleware.Recover())

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	return r
}
