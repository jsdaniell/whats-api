// Middlewares to handle generic handles of request.
package middlewares

import (
	"fmt"
	"github.com/jsdaniell/whats_api/api/responses"
	"log"
	"net/http"
)

// Validate Access Control to API and deal with CORS request.
func SetMiddlewareLogger(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		// TODO: When in production setup to https://website

		w.Header().Set("Access-Control-Allow-Origin", "*")

		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Methods", "*")

		log.Println("%s %s%s %s", r.Method, r.Host, r.RequestURI, r.Proto)

		if (*r).Method == "OPTIONS" {
			return
		}

		next(w, r)
	}
}

// Middleware to test authorization header parameter on closed routes.
func SetMiddlewareJSON(next http.HandlerFunc, openRoute bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if !openRoute {
			auth := r.Header.Get("Authorization")
			if auth == "" {
				responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("missing authorization token"))
				return
			}
		}

		next(w, r)
	}
}
