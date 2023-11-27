package main

import (
	"github.com/noam-g4/go.lib/easyblog/handlers/home"
	"github.com/noam-g4/go.lib/easyblog/handlers/post"
	"github.com/noam-g4/go.lib/easyblog/middlewares/auth"
	"github.com/noam-g4/go.lib/easyblog/server"
	"go.uber.org/fx"
)

func main() {
	handlersGroup := fx.ResultTags(`group:"handlers"`)
	middlewaresGroup := fx.ResultTags(`group:"middlewares"`)

	fx.New(
		fx.Provide(
			server.NewServer,
			server.NewServeMux,

			// middlewares group
			fx.Annotate(auth.NewAuthMiddleware, middlewaresGroup),

			// handlers group
			fx.Annotate(home.NewHomeHandler, handlersGroup),
			fx.Annotate(post.NewPostHandler, handlersGroup),
		),

		fx.Invoke(server.Start),
	).Run()
}
