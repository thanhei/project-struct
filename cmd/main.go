package main

import (
	"context"
	"fmt"
	"go-training/internal/app"
	"go-training/internal/common"
	"go-training/internal/pkg/subscriber"
	"log"
	"net/http"

	"go-training/internal/component/pubsub"
	restaurantrepo "go-training/internal/modules/restaurant/repository"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func main() {
	app.New(
		fx.Invoke(run),
		fx.Invoke(runSubscriber),
	).Run()
}

type RunParams struct {
	fx.In

	Lc      fx.Lifecycle
	Cfg     *common.Config
	Router  *gin.Engine
	Modules []common.AppModule `group:"modules"`
}

func run(p RunParams) {
	var srv *http.Server

	p.Lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			cfg := p.Cfg
			r := p.Router

			for _, m := range p.Modules {
				m.SetupRoutes(r.Group("/v1"))
			}

			srv = &http.Server{
				Addr:    fmt.Sprintf(":%d", cfg.Service.Port),
				Handler: r,
			}

			go func() {
				if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					log.Fatalf("listen: %s\n", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			if err := srv.Shutdown(ctx); err != nil {
				log.Fatal("Server forced to shutdown: ", err)
			}
			log.Println("Server exiting")
			return nil
		},
	})
}

func runSubscriber(lc fx.Lifecycle, pubsub pubsub.Pubsub, restaurantRepo restaurantrepo.RestaurantRepository) {
	e := subscriber.NewEngine(pubsub, restaurantRepo)
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Println("Start subscriber")
			return e.Start()
		},
	})
}
