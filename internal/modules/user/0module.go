package user

import (
	"go-training/internal/common"
	userbiz "go-training/internal/modules/user/business"
	userrepo "go-training/internal/modules/user/repository"
	userreposql "go-training/internal/modules/user/repository/sql"
	userapi "go-training/internal/modules/user/transport/api"

	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module(
		"user",
		fx.Provide(
			fx.Annotate(
				userreposql.NewSQLRepo,
				fx.As(new(userrepo.UserRepository)),
			),
			fx.Annotate(
				userbiz.NewBusiness,
				fx.As(new(userbiz.UserBusiness)),
			),
			common.AsAppModule(userapi.NewApi),
		),
	)
}
