package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const pathPrefix = "/api"

// GetRouter obtaining router
func (a API) GetRouter() *mux.Router {
	fmt.Println("Creating the router...")
	router := mux.NewRouter()
	s := router.PathPrefix(pathPrefix).Subrouter()

	s.HandleFunc("/hello", HTTPHandler(a.helloWorld)).Methods(http.MethodGet)

	return s
}