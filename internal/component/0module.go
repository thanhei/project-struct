package component

import (
	"go-training/internal/component/database"
	"go-training/internal/component/hasher"
	"go-training/internal/component/hasher/md5"
	"go-training/internal/component/pubsub"
	"go-training/internal/component/pubsub/pblocal"
	"go-training/internal/component/tokenprovider"
	"go-training/internal/component/tokenprovider/jwt"

	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("component",
		fx.Provide(
			database.New,
			fx.Annotate(pblocal.NewPubSub, fx.As(new(pubsub.Pubsub))),
			fx.Annotate(md5.NewMd5Hash, fx.As(new(hasher.Hasher))),
			fx.Annotate(
				jwt.NewTokenJWTProvider,
				fx.As(new(tokenprovider.Provider)),
			),
		),
	)
}
