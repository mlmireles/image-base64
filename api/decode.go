package api

import (
	"net/http"

	e "github.com/mlmireles/image-base64/errors"
)

func (a API) decode(w http.ResponseWriter, r *http.Request) e.HTTPError {

	return respond(w, nil)
}
