package rest

import (
	"net/http"

	"github.com/go-chi/chi"
)

type Mux struct {
	*chi.Mux
}

func NewRouter() *Mux {
	return &Mux{
		chi.NewRouter(),
	}
}

func (r *Mux) MountHandler(pattern string, handler interface{}) {
	restHandler, ok := handler.(Handler)
	if !ok {
		panic("expect a handler that implements `rest.Handler` interface")
	}
	restHandler.setChild(restHandler)

	httpHandler, ok := handler.(http.Handler)
	if !ok {
		panic("expect a handler that implements `http.Handler` interface")
	}

	r.Mount(pattern, httpHandler)
}
