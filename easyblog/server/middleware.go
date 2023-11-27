package server

import (
	"net/http"
)

type Middleware func(http.ResponseWriter, *http.Request) error

type middlewareChain struct {
	middleware Middleware
	next       *middlewareChain
}

type wrappedHandler struct {
	handleFunc http.HandlerFunc
}

func attachMiddlewares(middlewares []Middleware, handler Handler) http.Handler {
	if middlewares == nil {
		return handler
	}

	chain := newMiddlewareChain()
	chain.buildChain(middlewares)

	wrapper := func(w http.ResponseWriter, r *http.Request) {
		mw := chain
		for mw.nextMiddleware() {
			if err := mw.middleware(w, r); err != nil {
				return
			}
			mw = mw.next
		}

		handler.ServeHTTP(w, r)
	}

	return &wrappedHandler{handleFunc: wrapper}
}

func newMiddlewareChain() *middlewareChain {
	c := middlewareChain{}
	return &c
}

func (chain *middlewareChain) buildChain(mws []Middleware) {
	if len(mws) == 0 {
		chain.middleware = func(http.ResponseWriter, *http.Request) error { return nil }
		return
	}

	chain.middleware = mws[0]
	next := middlewareChain{}

	chain.next = &next
	next.buildChain(mws[1:])
}

func (chain *middlewareChain) nextMiddleware() bool {
	return chain.next != nil
}

func (wrapper *wrappedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	wrapper.handleFunc(w, r)
}
