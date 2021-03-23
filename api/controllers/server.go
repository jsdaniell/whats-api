package controllers

import "net/http"

// Handle the root / route to return feedback about the server to request like "Server Running..."
func ServerRunning(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server Running..."))
}
