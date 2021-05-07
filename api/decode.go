package api

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"

	e "github.com/mlmireles/image-base64/errors"
)

func (a API) decode(w http.ResponseWriter, r *http.Request) e.HTTPError {
	var encoded string
	err := json.NewDecoder(r.Body).Decode(&encoded)
	if err != nil {
		log.Println("[encode.go] ", err)
		return e.HTTPError{Error: e.BadRequest{}, Message: "Error reading file"}
	}

	b, err := fromBase64(encoded)
	if err != nil {
		log.Println("[encode.go] ", err)
		return e.HTTPError{Error: e.BadRequest{}, Message: "Error reading file"}
	}

	return respond(w, b)
}

func fromBase64(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}
