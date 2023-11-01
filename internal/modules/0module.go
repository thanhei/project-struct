package modules

import (
	"go-training/internal/modules/restaurant"
	"go-training/internal/modules/restaurantlike"
	"go-training/internal/modules/user"

	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module(
		"modules",
		user.Module(),
		restaurant.Module(),
		restaurantlike.Module(),
	)
}
