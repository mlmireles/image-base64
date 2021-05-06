package main

import (
	"fmt"
	"net/http"

	"image-base64/api"
	"image-base64/config"
)

func main() {
	fmt.Println("Image to Base64")

	a := api.API{}
	r := a.GetRouter()
	port, h := config.GetServerHandler(r)
	http.ListenAndServe(port, h)
}
