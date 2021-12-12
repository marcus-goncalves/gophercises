package api

import (
	"fmt"
	"html/template"
	"net/http"
)

var (
	Routes ssRoutes = &SsRoutes{}
	err    error
)

type SsRoutes struct{}

type ssRoutes interface {
	Home(w http.ResponseWriter, r *http.Request)
	Contact(w http.ResponseWriter, r *http.Request)
	Faq(w http.ResponseWriter, r *http.Request)
	NotFoundPage(w http.ResponseWriter, r *http.Request)
}

type View struct {
	Template *template.Template
	Layout   string
}

func NewView(layout string, files ...string) *View {
	var t *template.Template
	files = append(files,
		"server_side_app/views/layouts/footer.html",
		"server_side_app/views/layouts/main.html")
	t, err = template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
		Layout:   layout,
	}
}

func (route *SsRoutes) Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	// Route Logic
	homeView := NewView("content", "server_side_app/views/home.html")

	err = homeView.Template.ExecuteTemplate(w, homeView.Layout, nil)
	if err != nil {
		panic(err)
	}

}

func (route *SsRoutes) Contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	// Route Logic
	contactView := NewView("content", "server_side_app/views/contact.html")

	err = contactView.Template.ExecuteTemplate(w, contactView.Layout, nil)
	if err != nil {
		panic(err)
	}

}

func (route *SsRoutes) Faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Hi there! FAQ Page</h1>")

}

func (route *SsRoutes) NotFoundPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>Hi there! Not Found!</h1>")

}
