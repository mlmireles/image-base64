// +build dev

package config

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func getEnv() string {
	return "env-dev"
}

// GetServerHandler obtain http.Handler with CORS
func GetServerHandler(r *mux.Router) (string, http.Handler) {
	fmt.Println("Starting server at port :4040")
	return ":4040", getCORS().Handler(r)
}

func getCORS() *cors.Cors {
	// (Develop ONLY) Create the Server allow CORS
	// REMOVE AT PRODUCTION
	fmt.Println("WARNING! ------ CORS allow ------ Development only")
	return cors.AllowAll()
}
