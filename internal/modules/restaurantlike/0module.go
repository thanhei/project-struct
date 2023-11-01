package restaurantlike

import (
	"go-training/internal/common"
	restaurantlikebiz "go-training/internal/modules/restaurantlike/business"
	restaurantlikerepo "go-training/internal/modules/restaurantlike/repository"
	restaurantlikereposql "go-training/internal/modules/restaurantlike/repository/sql"
	restaurantlikeapi "go-training/internal/modules/restaurantlike/transport/api"

	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module(
		"restaurantlike",
		fx.Provide(
			fx.Annotate(
				restaurantlikereposql.NewSQLRepo,
				fx.As(new(restaurantlikerepo.RestaurantLikeRepository)),
			),
			fx.Annotate(
				restaurantlikebiz.NewBusiness,
				fx.As(new(restaurantlikebiz.RestaurantLikeBusiness)),
			),
			common.AsAppModule(restaurantlikeapi.NewApi),
		),
	)
}
