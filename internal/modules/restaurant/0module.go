package restaurant

import (
	"go-training/internal/common"
	restaurantbiz "go-training/internal/modules/restaurant/business"
	restaurantrepo "go-training/internal/modules/restaurant/repository"
	restaurantreposql "go-training/internal/modules/restaurant/repository/sql"
	restaurantapi "go-training/internal/modules/restaurant/transport/api"

	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module(
		"restaurant",
		fx.Provide(
			fx.Annotate(
				restaurantreposql.NewSQLRepo,
				fx.As(new(restaurantrepo.RestaurantRepository)),
			),
			fx.Annotate(
				restaurantbiz.NewBusiness,
				fx.As(new(restaurantbiz.RestaurantBusiness)),
			),
			common.AsAppModule(restaurantapi.NewApi),
		),
	)
}
