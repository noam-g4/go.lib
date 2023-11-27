package server

import (
	"net/http"

	"go.uber.org/fx"
)

type Handler interface {
	http.Handler
	Pattern() string
}

type muxDeps struct {
	fx.In
	Handlers    []Handler    `group:"handlers"`
	Middlewares []Middleware `group:"middlewares"`
}

func NewServeMux(deps muxDeps) *http.ServeMux {
	m := http.NewServeMux()

	fs := http.FileServer(http.Dir("./static"))
	m.Handle("/static/", http.StripPrefix("/static/", fs))

	for _, h := range deps.Handlers {
		m.Handle(h.Pattern(), attachMiddlewares(deps.Middlewares, h))
	}

	return m
}
