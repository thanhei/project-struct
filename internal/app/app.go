package app

import (
	"go-training/internal/common"
	"go-training/internal/component"
	"go-training/internal/middleware"
	"go-training/internal/modules"

	"go-training/internal/server"

	"go.uber.org/fx"
)

func Options(additionalOpts ...fx.Option) []fx.Option {
	cfg, err := common.LoadConfig()
	if err != nil {
		panic(err)
	}
	baseOpts := []fx.Option{
		fx.Supply(cfg),
		component.Module(),
		middleware.Module(),
		modules.Module(),
		fx.Provide(
			server.NewServer,
		),
	}

	return append(baseOpts, additionalOpts...)
}

func New(additionalOpts ...fx.Option) *fx.App {
	return fx.New(
		Options(additionalOpts...)...,
	)
}
