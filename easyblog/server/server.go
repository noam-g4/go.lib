package server

import (
	"context"
	"net"
	"net/http"

	"go.uber.org/fx"
)

type deps struct {
	fx.In

	LC      fx.Lifecycle
	Handler *http.ServeMux
}

func NewServer(d deps) *http.Server {
	// TODO: add config to fetch port
	srv := http.Server{Addr: ":2508", Handler: d.Handler}

	d.LC.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}

			// TODO: add logger to log the startup
			go srv.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})

	return &srv
}

func Start(*http.Server) {}
