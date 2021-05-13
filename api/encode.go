package api

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

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

	var bytes []byte
	if !strings.HasPrefix(encode.Path, "http") {
		// input file directory
		// Read the entire file into a byte slice
		bytes, err = ioutil.ReadFile(encode.Path)
		if err != nil {
			log.Println("[encode.go] ", err)
			return e.HTTPError{Error: e.BadRequest{}, Message: "Error reading file"}
		}

	} else {
		// input http path
		resp, err := http.Get(encode.Path)
		if err != nil {
			log.Println("[encode.go] ", err)
			return e.HTTPError{Error: e.BadRequest{}, Message: "Error getting file from url"}
		}

		defer resp.Body.Close()
		bytes, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
			return e.HTTPError{Error: e.BadRequest{}, Message: "Error reading file from url"}
		}
	}

	var encoded string
	mimeType := http.DetectContentType(bytes)

	switch mimeType {
	case "image/jpeg":
		encoded += "data:image/jpeg;base64,"
	case "image/png":
		encoded += "data:image/png;base64,"
	}

	encoded += toBase64(bytes)
	//log.Println(encoded)
	return respond(w, encoded)
}

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
