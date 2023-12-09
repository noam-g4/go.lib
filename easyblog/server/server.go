package server

import (
	"context"
	"log"
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
	srv := http.Server{Addr: ":8080", Handler: d.Handler}

	d.LC.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}

			log.Print("server is starting")
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
