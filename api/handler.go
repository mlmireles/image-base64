package api

import (
	"encoding/json"
	"log"
	"net/http"

	e "github.com/mlmireles/image-base64/errors"
)

// HTTPHandler wraps a function returning an error by handling the error and
// returning a http.Handler.
// If the error is of the one of the types defined above, it is handled as
// described for every type.
// If the error is of another type, it is considered as an internal error and
// its message is logged.
func HTTPHandler(f func(w http.ResponseWriter, r *http.Request) e.HTTPError) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err.Error == nil {
			log.Printf("(%d) %s: %s", 200, r.Method, r.URL)
			return
		}

		message, _ := json.Marshal(err)

		var status int
		switch err.Error.(type) {
		case e.BadRequest:
			status = http.StatusBadRequest
			http.Error(w, string(message), http.StatusBadRequest)
			break
		case e.NotFound:
			status = http.StatusNotFound
			http.Error(w, string(message), http.StatusNotFound)
			break
		case e.NotAuthorized:
			status = http.StatusUnauthorized
			http.Error(w, string(message), http.StatusUnauthorized)
			break
		case e.UnprocessableEntity:
			status = http.StatusUnprocessableEntity
			http.Error(w, string(message), http.StatusUnprocessableEntity)
			break
		case e.DuplicateItem:
			status = http.StatusNotAcceptable
			http.Error(w, string(message), http.StatusNotAcceptable)
			break
		default:
			status = http.StatusInternalServerError
			http.Error(w, err.Message, http.StatusInternalServerError)
			break
		}

		log.Printf("[Error] (%d) %s %s \n%s", status, r.Method, r.URL, err.Message)
	}
}
