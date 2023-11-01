package common

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type AppModule interface {
	SetupRoutes(r *gin.RouterGroup)
}

func AsAppModule(f any) any {
	return fx.Annotate(f,
		fx.As(new(AppModule)),
		fx.ResultTags(`group:"modules"`),
	)
}
