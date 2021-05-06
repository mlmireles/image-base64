package api

import (
	"encoding/json"
	"log"
	"net/http"

	e "image-base64/errors"
)

// API struct interface
type API struct{}

func (a API) helloWorld(w http.ResponseWriter, r *http.Request) e.HTTPError {
	log.Println("Hello World!")

	return respond(w, "Hello World!")
}

func respond(w http.ResponseWriter, payload interface{}) e.HTTPError {
	resp, err := json.Marshal(payload)
	if err != nil {
		log.Println(err.Error())
		return e.HTTPError{Error: e.BadRequest{}}
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)

	return e.HTTPError{}
}
