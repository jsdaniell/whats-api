package api

import (
	"fmt"
	"github.com/jsdaniell/whats_api/api/router"
	"github.com/jsdaniell/whats_api/config"
	"log"
	"net/http"
)

// Run initialize the server.
func Run() {
	config.Load()
	fmt.Printf("Listening... localhost:%d", config.PORT)
	listen(config.PORT)
}

// Configure listen port.
func listen(port int) {
	r := router.New()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
