// +build live

package config

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// GetServerHandler Obtains http.Handler
func GetServerHandler(r *mux.Router) (string, http.Handler) {
	fmt.Println("Starting server at port :4040")
	return ":" + "4040", r
}
