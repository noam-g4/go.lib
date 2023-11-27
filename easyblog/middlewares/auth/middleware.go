package auth

import (
	"net/http"

	"github.com/noam-g4/go.lib/easyblog/server"
	"go.uber.org/fx"
)

// TODO: implement middleware
type authMiddlewareDeps struct {
	fx.In
}

func NewAuthMiddleware(deps authMiddlewareDeps) server.Middleware {
	return func(w http.ResponseWriter, r *http.Request) error {
		// TODO: implement
		return nil
	}
}
