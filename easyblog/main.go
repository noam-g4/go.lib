package main

import (
	"github.com/noam-g4/go.lib/easyblog/handler"
	"github.com/noam-g4/go.lib/easyblog/server"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			server.NewServer,
			server.NewServeMux,
			handler.NewHandlersGenerator,
			handler.NewTemplates,
		),

		fx.Invoke(server.Start),
		// fx.NopLogger,
	).Run()
}
