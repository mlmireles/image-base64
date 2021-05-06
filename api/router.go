package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const pathPrefix = ""

// GetRouter obtaining router
func (a API) GetRouter() *mux.Router {
	fmt.Println("Creating router...")
	router := mux.NewRouter()
	s := router.PathPrefix(pathPrefix).Subrouter()

	s.HandleFunc("/hello", HTTPHandler(a.helloWorld)).Methods(http.MethodGet)
	s.HandleFunc("/encode", HTTPHandler(a.encode)).Methods(http.MethodGet)

	return s
}
