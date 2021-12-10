package main

import (
	"log"
	"net/http"

	api "photo_gallery.com/api/apis"
)

func main() {
	r := api.GorillaRoutes()
	// r := api.ChiRoutes()

	log.Fatal(http.ListenAndServe(":3000", r))
}
