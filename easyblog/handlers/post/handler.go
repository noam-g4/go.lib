package post

import (
	"net/http"

	"github.com/noam-g4/go.lib/easyblog/server"
	"go.uber.org/fx"
)

// TODO: implement handler
type handler struct{}

type deps struct {
	fx.In
}

func NewPostHandler(deps deps) server.Handler {
	return &handler{}
}

func (h *handler) Pattern() string {
	return "/post"
}

func (h *handler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	// TODO: implement
	w.Write([]byte(h.Pattern()))
}
