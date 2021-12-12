package api

import (
	"fmt"
	"net/http"
)

var (
	Routes apiRoutes = &ApiRoutes{}
)

type ApiRoutes struct{}

type apiRoutes interface {
	Home(w http.ResponseWriter, r *http.Request)
	Contact(w http.ResponseWriter, r *http.Request)
	Faq(w http.ResponseWriter, r *http.Request)
	NotFoundPage(w http.ResponseWriter, r *http.Request)
}

func (route *ApiRoutes) Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Hi there! Welcome Page</h1>")
}

func (route *ApiRoutes) Contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Hi there! Contact Page</h1>")

}

func (route *ApiRoutes) Faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Hi there! FAQ Page</h1>")

}

func (route *ApiRoutes) NotFoundPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>Hi there! Not Found!</h1>")

}
