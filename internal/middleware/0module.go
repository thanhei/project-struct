package middleware

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Module("middleware",
		fx.Provide(
			fx.Annotate(
				RequireAuth,
				fx.ResultTags(`name:"requireAuth"`),
			),
		),
	)
}
