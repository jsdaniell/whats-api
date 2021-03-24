package routes

import (
	"github.com/jsdaniell/whats_api/api/controllers"
	"net/http"
)

var authenticationRoutes = []Route{
	{
		URI:     "/getQrCode/{sessionID}",
		Method:  http.MethodGet,
		Handler: controllers.Connect,
		Open:    true,
	},
	{
		URI:     "/disconnect/{sessionID}",
		Method:  http.MethodGet,
		Handler: controllers.Disconnect,
		Open:    true,
	},
}
