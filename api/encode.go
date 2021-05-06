package api

import (
	"encoding/json"
	"log"
	"net/http"

	e "github.com/mlmireles/image-base64/errors"
	"github.com/mlmireles/image-base64/models"
)

func (a API) encode(w http.ResponseWriter, r *http.Request) e.HTTPError {
	var encode models.Encode
	err := json.NewDecoder(r.Body).Decode(&encode)
	if err != nil {
		log.Println("[encode.go]", err.Error())
		return e.HTTPError{Error: e.BadRequest{}, Message: "Error reading body"}
	}

	return respond(w, encode)
}
