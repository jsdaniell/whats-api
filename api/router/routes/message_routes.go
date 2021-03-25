package routes

import (
	"github.com/jsdaniell/whats_api/api/controllers"
	"net/http"
)

var messageRoutes = []Route{
	{
		URI:     "/sendMessage/{sessionID}",
		Method:  http.MethodPost,
		Handler: controllers.Message,
		Open:    true,
	},
}
