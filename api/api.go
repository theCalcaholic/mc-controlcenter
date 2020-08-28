package api

import (
	"net/http"
	"regexp"
)

type route struct {
	pattern *regexp.Regexp
	handler http.Handler
}

// RegexpHandler all api routes
type RegexpHandler struct {
	routes []*route
}

// Handler Add handler
func (h *RegexpHandler) Handler(pattern *regexp.Regexp, handler http.Handler) {
	h.routes = append(h.routes, &route{pattern, handler})
}

// HandleFunc add handler from function
func (h *RegexpHandler) HandlerFunc(pattern *regexp.Regexp, handlerFn func(http.ResponseWriter, *http.Request)) {
	h.routes = append(h.routes, &route{pattern, http.HandlerFunc(handlerFn)})
}

// Serve Register handlers for routes
func (h *RegexpHandler) Serve(w http.ResponseWriter, r *http.Request) {

	http.NotFound(w, r)
}

// API get api handlers
func API() func(http.ResponseWriter, *http.Request) {
	var handlers = new(RegexpHandler)
	pattern := regexp.MustCompile("/api/server/(?P<server-name>.+)/stop")
	handlers.HandlerFunc(pattern, serverStop)

	return handlers.Serve
}

func serverStop(w http.ResponseWriter, r *http.Request) {

}
