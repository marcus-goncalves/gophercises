package api

// func ChiRoutes() *chi.Mux {
// 	r := chi.NewRouter()

// 	// Middlewares
// 	r.Use(middleware.Logger)

// 	// Routes
// 	r.Get("/", home)
// 	r.Get("/contact", contact)
// 	r.Get("/faq", faq)
// 	r.NotFound(notFound)

// 	return r
// }

// func home(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html")
// 	fmt.Fprint(w, "<h1>Hi there! Welcome Page</h1>")
// }

// func contact(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html")
// 	fmt.Fprint(w, "<h1>Hi there! Contact Page</h1>")

// }

// func faq(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html")
// 	fmt.Fprint(w, "<h1>Hi there! FAQ Page</h1>")

// }

// func notFound(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html")
// 	w.WriteHeader(http.StatusNotFound)
// 	fmt.Fprint(w, "<h1>Hi there! Not Found!</h1>")

// }
