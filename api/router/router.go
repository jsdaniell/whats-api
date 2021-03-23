// Router setup.
package router

import (
	"github.com/gorilla/mux"
	"github.com/jsdaniell/whats_api/api/router/routes"
)

func New() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	return routes.SetupRoutesWithMiddlewares(r)
}
