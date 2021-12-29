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
	r.Handle("/", controllers.NewStatic().HomeView).Methods("GET")
	r.Handle("/contact", controllers.NewStatic().ContactView).Methods("GET")
	r.Handle("/faq", controllers.NewStatic().FaqView).Methods("GET")

	r.HandleFunc("/signup", controllers.NewUsers().New).Methods("GET")
	r.HandleFunc("/signup", controllers.NewUsers().Create).Methods("POST")

	r.NotFoundHandler = http.HandlerFunc(api.Routes.NotFoundPage)

	log.Fatal(http.ListenAndServe(":3000", r))

}
