package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	api "photo_gallery.com/v1/server_side_app"
	"photo_gallery.com/v1/server_side_app/controllers"
)

// func main() {
// 	// r := api.GorillaRoutes()
// 	r := chi.NewRouter()

// 	// Routes
// 	r.Get("/", api.Routes.Home)
// 	r.Get("/contacts", api.Routes.Contact)
// 	r.Get("/faq", api.Routes.Faq)
// 	r.NotFound(api.Routes.NotFoundPage)

// 	log.Fatal(http.ListenAndServe(":3000", r))
// }

func main() {
	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/", api.Routes.Home)
	r.HandleFunc("/contact", api.Routes.Contact)
	r.HandleFunc("/faq", api.Routes.Faq)
	r.HandleFunc("/signup", controllers.NewUsers().New)
	r.NotFoundHandler = http.HandlerFunc(api.Routes.NotFoundPage)

	log.Fatal(http.ListenAndServe(":3000", r))

}
