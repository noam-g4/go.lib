package server

import (
	"net/http"

	"github.com/noam-g4/go.lib/easyblog/handler"
	"go.uber.org/fx"
)

type muxDeps struct {
	fx.In
	HandlerFactory *handler.HandlersGenerator
}

func NewServeMux(deps muxDeps) *http.ServeMux {
	m := http.NewServeMux()

	fs := http.FileServer(http.Dir("./content/static"))
	m.Handle("/content/static/", http.StripPrefix("/content/static/", fs))

	for _, h := range deps.HandlerFactory.Handlers() {
		m.Handle(h.Pattern(), h)
	}

	return m
}
