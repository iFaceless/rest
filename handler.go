package rest

import (
	"net/http"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/spf13/cast"
)

type BaseHandler struct {
	childHandler Handler
	W            http.ResponseWriter
	R            *http.Request
}

func (hd *BaseHandler) setChild(childHandler Handler) {
	hd.childHandler = childHandler
}

// ServeHTTP implements the `http.server.Handler` interface
func (hd *BaseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hd.W = w
	hd.R = r

	var handler Handler
	if hd.childHandler != nil {
		handler = hd.childHandler
	} else {
		handler = hd
	}

	// make sure `Finish` is called
	defer handler.Finish()
	handler.Prepare()

	switch r.Method {
	case http.MethodGet:
		handler.Get()
	case http.MethodPost:
		handler.Post()
	case http.MethodPut:
		handler.Put()
	case http.MethodPatch:
		handler.Patch()
	case http.MethodDelete:
		handler.Delete()
	default:
		hd.RenderError(HTTPMethodNotAllowedError)
	}
}

func (hd *BaseHandler) Prepare() {

}

// Finish must be called after serving a request
func (hd *BaseHandler) Finish() {

}

func (hd *BaseHandler) Get() {
	hd.RenderError(HTTPMethodNotAllowedError)
}

func (hd *BaseHandler) Post() {
	hd.RenderError(HTTPMethodNotAllowedError)
}

func (hd *BaseHandler) Patch() {
	hd.RenderError(HTTPMethodNotAllowedError)
}

func (hd *BaseHandler) Put() {
	hd.RenderError(HTTPMethodNotAllowedError)
}

func (hd *BaseHandler) Delete() {
	hd.RenderError(HTTPMethodNotAllowedError)
}

func (hd *BaseHandler) RenderJSON(v interface{}) {
	hd.W.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(hd.W).Encode(v); err != nil {
		http.Error(hd.W, err.Error(), http.StatusInternalServerError)
	}
}

func (hd *BaseHandler) RenderError(err error) {
	apiError, ok := err.(*APIError)
	if ok {
		if apiError.HTTPStatusCode >= http.StatusContinue && apiError.HTTPStatusCode <= http.StatusNetworkAuthenticationRequired {
			hd.W.WriteHeader(apiError.HTTPStatusCode)
		} else {
			hd.W.WriteHeader(http.StatusBadRequest)
		}
	}
	hd.RenderJSON(apiError)
}

func (hd *BaseHandler) Offset() int {
	return hd.IntQueryArgument("offset", defaultOffset)
}

func (hd *BaseHandler) Limit() int {
	return hd.IntQueryArgument("limit", defaultLimit)
}

func (hd *BaseHandler) URLParam(key string) string {
	return chi.URLParam(hd.R, key)
}

func (hd *BaseHandler) QueryArgument(key string) string {
	return hd.R.URL.Query().Get(key)
}

func (hd *BaseHandler) IntQueryArgument(key string, fallback int) int {
	if value, err := cast.ToIntE(hd.QueryArgument(key)); err != nil {
		return fallback
	} else {
		return value
	}
}

func (hd *BaseHandler) Int64QueryArgument(key string, fallback int64) int64 {
	if value, err := cast.ToInt64E(hd.QueryArgument(key)); err != nil {
		return fallback
	} else {
		return value
	}
}
