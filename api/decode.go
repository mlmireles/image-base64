package api

import (
	"bytes"
	"encoding/base64"
	"log"
	"net/http"
	"strings"

	e "github.com/mlmireles/image-base64/errors"
)

func (a API) decode(w http.ResponseWriter, r *http.Request) e.HTTPError {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	encoded := buf.String()

	arr := strings.Split(encoded, ",")
	b, err := fromBase64(arr[1])
	if err != nil {
		log.Println("[encode.go] ", err)
		return e.HTTPError{Error: e.BadRequest{}, Message: "Error reading file"}
	}

	return respond(w, b)
}

func fromBase64(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}
