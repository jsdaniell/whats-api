package routes

import (
	"github.com/jsdaniell/whats_api/api/controllers"
	"net/http"
)

var authenticationRoutes = []Route{
	{
		URI:     "/getQrCode",
		Method:  http.MethodGet,
		Handler: controllers.Connect,
		Open:    true,
	},
	{
		URI:     "/disconnect",
		Method:  http.MethodGet,
		Handler: controllers.Disconnect,
		Open:    true,
	},
}
