package routes

import (
	"github.com/jsdaniell/whats_api/api/controllers"
	"net/http"
)

// Generic server routes with their controller.
var serverRoutes = []Route{
	{
		URI:     "/",
		Method:  http.MethodGet,
		Handler: controllers.ServerRunning,
		Open:    true,
	},
}
