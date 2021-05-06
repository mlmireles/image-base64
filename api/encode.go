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
			log.Fatal(err)
		}

	} else {
		// input http path
		resp, err := http.Get(encode.Path)
		if err != nil {
			log.Fatal(err)
		}

		defer resp.Body.Close()
		bytes, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
	}

	var base64Encoding string
	mimeType := http.DetectContentType(bytes)

	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	}

	base64Encoding += toBase64(bytes)
	log.Println(base64Encoding)
	return respond(w, encode)
}

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
