package routes

import (
	"github.com/gorilla/mux"
	"github.com/jsdaniell/whats_api/api/middlewares"
	"net/http"
)

// Route struct provides the structure of route
type Route struct {
	URI     string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
	Open    bool
}

// Load all routes and setting with middlewares to receive requests.
func Load() []Route {
	routes := [][]Route{
		serverRoutes,
		authenticationRoutes,
		messageRoutes,
	}

	var joinedRoutes []Route

	for _, r := range routes {
		joinedRoutes = append(joinedRoutes, r...)
	}

	return joinedRoutes
}

// SetupRoutesWithMiddlewares Using middlewares on handle of route.
func SetupRoutesWithMiddlewares(r *mux.Router) *mux.Router {

	for _, route := range Load() {
		r.HandleFunc(route.URI, middlewares.SetMiddlewareLogger(
			middlewares.SetMiddlewareJSON(route.Handler, route.Open))).Methods(route.Method, "OPTIONS")
	}

	return r
}
